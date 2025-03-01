package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martishin/movie-search-service/internal/handler"
	"github.com/martishin/movie-search-service/internal/model/config"
	"github.com/martishin/movie-search-service/internal/repository"
	"github.com/martishin/movie-search-service/internal/route"
	"github.com/martishin/movie-search-service/internal/service"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func configureGoogleOauth(config *config.OAuthConfig) {
	store := sessions.NewCookieStore([]byte(config.SessionSecret))

	store.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   config.IsProduction, // Enable secure cookies in production
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // 30 days
		Domain:   config.Domain,
	}

	gothic.Store = store //nolint:reassign

	// Configure Google provider
	goth.UseProviders(
		google.New(config.ClientID, config.ClientSecret, config.CallbackURL, "email", "profile"),
	)
}

func NewServer(logger *slog.Logger, pool *pgxpool.Pool, serverConfig *config.ServerConfig, oauthConfig *config.OAuthConfig) *http.Server {
	// Initialize repositories
	userRepo := repository.NewUserRepository(pool)

	// Initialise services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService, oauthConfig)

	// Create Server instance
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConfig.Port),
		Handler:      route.RegisterRoutes(logger, userHandler, authHandler),
		IdleTimeout:  serverConfig.IdleTimeout,
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
	}

	logger.Info("Server port", slog.Int("port", serverConfig.Port))
	logger.Info("Cookie domain", slog.String("domain", oauthConfig.Domain))
	logger.Info("Environment mode", slog.Bool("is_production", oauthConfig.IsProduction))

	// Configure OAuth
	configureGoogleOauth(oauthConfig)
	logger.Info("Google OAuth provider configured", slog.String("callback_url", oauthConfig.CallbackURL))

	return server
}
