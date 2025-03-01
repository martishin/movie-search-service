package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratepgx "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// RunPostgresMigrations applies database migrations
func RunPostgresMigrations(pool *pgxpool.Pool) error {
	driver, err := migratepgx.WithInstance(stdlib.OpenDBFromPool(pool), &migratepgx.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}
	defer func() {
		if closeErr := driver.Close(); closeErr != nil {
			fmt.Printf("Failed to close migration driver: %v\n", closeErr)
		}
	}()

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
