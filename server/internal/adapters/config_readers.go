package adapters

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/martishin/movie-search-service/internal/models"

	_ "github.com/joho/godotenv/autoload"
)

func ReadServerConfig() (*models.ServerConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid or missing PORT environment variable")
	}

	return &models.ServerConfig{
		Port:         port,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}

func ReadGoogleOauthConfig() (*models.OAuthConfig, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	sessionSecret := os.Getenv("SESSION_SECRET")
	domain := os.Getenv("SESSION_COOKIE_DOMAIN")
	environment := os.Getenv("ENV")

	if clientID == "" || clientSecret == "" || callbackURL == "" ||
		sessionSecret == "" || domain == "" || environment == "" {
		return nil, fmt.Errorf("google OAuth environment variables are not set")
	}

	isProduction := environment == "production"

	return &models.OAuthConfig{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		CallbackURL:   callbackURL,
		SessionSecret: sessionSecret,
		Domain:        domain,
		IsProduction:  isProduction,
	}, nil
}
