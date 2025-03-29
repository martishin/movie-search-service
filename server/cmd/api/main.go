package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martishin/movie-search-service/internal/adapter"
	"github.com/martishin/movie-search-service/internal/db"
	"github.com/martishin/movie-search-service/internal/model/config"
	"github.com/martishin/movie-search-service/internal/server"
	"gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger(config *config.ObservabilityConfig) *slog.Logger {
	// Configure log rotation
	logFile := &lumberjack.Logger{
		Filename:   config.LogPath, // Rotated log file (relative path)
		MaxSize:    10,             // Max file size in MB before rotation
		MaxBackups: 24,             // Keep last 24 log files (1-day history)
		MaxAge:     1,              // Retain logs for 1 day
		Compress:   true,           // Compress rotated logs
	}

	// Create a multi-writer to log to both file and stdout
	multiWriter := slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), nil)

	// Initialize slog logger
	return slog.New(multiWriter)
}

func gracefulShutdown(logger *slog.Logger, apiServer *http.Server, pool *pgxpool.Pool, done chan struct{}) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()
	logger.Warn("Received shutdown signal. Initiating graceful shutdown...")

	// Create a context with a 5-second timeout for the server to finish requests
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Closing database connection pool...")
	pool.Close()

	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server forced to shutdown due to error", slog.Any("error", err))
	}

	// Notify that shutdown is complete
	close(done)
	logger.Info("Graceful shutdown complete. Exiting application.")
}

func main() {
	// Read Observability config
	observabilityConfig, err := adapter.ReadObservabilityConfig()
	if err != nil {
		os.Exit(1)
	}

	logger := setupLogger(observabilityConfig)

	// Read Postgres config
	postgresConfig, err := adapter.ReadPostgresConfig()
	if err != nil {
		logger.Error("Failed to read postgres config", slog.Any("error", err))
		os.Exit(1)
	}

	// Ensure database exists
	if err := db.EnsureDatabaseExists(postgresConfig); err != nil {
		logger.Error("Failed to ensure database exists", slog.Any("error", err))
		os.Exit(1)
	}

	// Connect to Postgres
	postgresPool, err := db.NewPostgresPool(postgresConfig)
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", slog.Any("error", err))
		os.Exit(1)
	}

	// Run Postgres migrations
	if err := db.RunPostgresMigrations(postgresPool); err != nil {
		logger.Error("Failed to apply migrations", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("Database migrations applied successfully")

	// Read Redis config
	redisConfig, err := adapter.ReadRedisConfig()
	if err != nil {
		logger.Error("Failed to read redis config", slog.Any("error", err))
		os.Exit(1)
	}

	// Connect to Redis
	redisClient, err := db.NewRedisClient(redisConfig)
	if err != nil {
		logger.Error("Failed to connect to Redis", slog.Any("error", err))
		os.Exit(1)
	}

	// Read server config
	serverConfig, err := adapter.ReadServerConfig()
	if err != nil {
		logger.Error("Failed to read server config", slog.Any("error", err))
		os.Exit(1)
	}

	// Read OAuth config
	oauthConfig, err := adapter.ReadGoogleOauthConfig()
	if err != nil {
		logger.Error("Failed to read Google OAuth config", slog.Any("error", err))
		os.Exit(1)
	}

	// Create the server
	serv := server.NewServer(
		logger,
		postgresPool,
		redisClient,
		serverConfig,
		oauthConfig,
		observabilityConfig,
	)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan struct{})

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(logger, serv, postgresPool, done)

	logger.Info("Starting server", slog.String("address", "http://localhost"+serv.Addr))
	err = serv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server error", slog.Any("error", err))
		os.Exit(1)
	}

	// Wait for the graceful shutdown to complete
	<-done
}
