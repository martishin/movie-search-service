package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martishin/movie-search-service/internal/adapters"
	"github.com/martishin/movie-search-service/internal/db"
	"github.com/martishin/movie-search-service/internal/server"
)

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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Read Postgres config
	postgresConfig, err := adapters.ReadPostgresConfig()
	if err != nil {
		logger.Error("Failed to read postgres config", slog.Any("error", err))
		os.Exit(1)
	}

	// Connect to Postgres
	pool, err := db.ConnectPostgresPool(postgresConfig)
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", slog.Any("error", err))
		os.Exit(1)
	}

	// Run Postgres migrations
	if err := db.RunPostgresMigrations(pool); err != nil {
		logger.Error("Failed to apply migrations", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("Database migrations applied successfully")

	// Read server config
	serverConfig, err := adapters.ReadServerConfig()
	if err != nil {
		logger.Error("Failed to read server config", slog.Any("error", err))
		os.Exit(1)
	}

	// Read OAuth config
	oauthConfig, err := adapters.ReadGoogleOauthConfig()
	if err != nil {
		logger.Error("Failed to read Google OAuth config", slog.Any("error", err))
		os.Exit(1)
	}

	// Create the server
	serv := server.NewServer(logger, serverConfig, oauthConfig)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan struct{})

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(logger, serv, pool, done)

	logger.Info("Starting server", slog.String("address", serv.Addr))
	err = serv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server error", slog.Any("error", err))
		os.Exit(1)
	}

	// Wait for the graceful shutdown to complete
	<-done
	logger.Info("Graceful shutdown complete.")
}
