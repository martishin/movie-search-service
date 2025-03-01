package adapter

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/martishin/movie-search-service/internal/model/config"

	_ "github.com/joho/godotenv/autoload"
)

func ReadServerConfig() (*config.ServerConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid or missing PORT environment variable")
	}

	return &config.ServerConfig{
		Port:         port,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}

func ReadGoogleOauthConfig() (*config.OAuthConfig, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	redirectURL := os.Getenv("REDIRECT_URL")
	sessionSecret := os.Getenv("SESSION_SECRET")
	domain := os.Getenv("SESSION_COOKIE_DOMAIN")
	environment := os.Getenv("ENV")

	if clientID == "" || clientSecret == "" || callbackURL == "" || redirectURL == "" ||
		sessionSecret == "" || domain == "" || environment == "" {
		return nil, fmt.Errorf("google OAuth environment variables are not set")
	}

	isProduction := environment == "production"

	return &config.OAuthConfig{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		CallbackURL:   callbackURL,
		RedirectURL:   redirectURL,
		SessionSecret: sessionSecret,
		Domain:        domain,
		IsProduction:  isProduction,
	}, nil
}

func ReadPostgresConfig() (*config.PostgresConfig, error) {
	host := os.Getenv("POSTGRES_HOST")
	portStr := os.Getenv("POSTGRES_PORT")
	database := os.Getenv("POSTGRES_DATABASE")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	if host == "" || portStr == "" || database == "" || username == "" || password == "" {
		return nil, fmt.Errorf("missing required PostgreSQL environment variables")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid POSTGRES_PORT: %v", err)
	}

	return &config.PostgresConfig{
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
	}, nil
}
