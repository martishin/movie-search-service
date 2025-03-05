package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martishin/movie-search-service/internal/model/config"
)

func NewPostgresPool(config *config.PostgresConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to reach PostgreSQL: %w", err)
	}

	return pool, nil
}

func EnsureDatabaseExists(config *config.PostgresConfig) error {
	// Connect to 'postgres' database first
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s/postgres",
		config.Username,
		config.Password,
		config.Host,
	)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	defer conn.Close(ctx)

	// Check if database exists
	var exists bool
	err = conn.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
		config.Database).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	// Create database if it doesn't exist
	if !exists {
		_, err = conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", config.Database))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
	}

	return nil
}
