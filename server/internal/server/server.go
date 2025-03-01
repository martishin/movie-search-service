package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/martishin/movie-search-service/internal/models/config"

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

func NewServer(logger *slog.Logger, serverConfig *config.ServerConfig, oauthConfig *config.OAuthConfig) *http.Server {
	// Create Server instance
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConfig.Port),
		Handler:      registerRoutes(logger),
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
