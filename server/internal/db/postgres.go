package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martishin/movie-search-service/internal/model/config"
)

func ConnectPostgresPool(config *config.PostgresConfig) (*pgxpool.Pool, error) {
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
