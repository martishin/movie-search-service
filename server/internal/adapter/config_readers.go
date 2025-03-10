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
	database := os.Getenv("POSTGRES_DATABASE")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	if host == "" || database == "" || username == "" || password == "" {
		return nil, fmt.Errorf("missing required PostgreSQL environment variables")
	}

	return &config.PostgresConfig{
		Host:     host,
		Database: database,
		Username: username,
		Password: password,
	}, nil
}

func ReadRedisConfig() (*config.RedisConfig, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	if host == "" || port == "" {
		return nil, fmt.Errorf("missing required Redis environment variables")
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_PORT: %v", err)
	}

	return &config.RedisConfig{
		Host: host,
		Port: port,
		DB:   db,
	}, nil
}

func ReadObservabilityConfig() (*config.ObservabilityConfig, error) {
	alloyUsername := os.Getenv("ALLOY_USERNAME")
	alloyPassword := os.Getenv("ALLOY_PASSWORD")
	logPath := os.Getenv("LOGS_PATH")

	if alloyUsername == "" || alloyPassword == "" || logPath == "" {
		return nil, fmt.Errorf("missing required observability environment variables")
	}

	return &config.ObservabilityConfig{
		AlloyUsername: alloyUsername,
		AlloyPassword: alloyPassword,
		LogPath:       logPath,
	}, nil
}
