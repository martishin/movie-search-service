package config

import "fmt"

type PostgresConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

// DSN returns the connection string for pgx
func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		c.Username, c.Password, c.Host, c.Database,
	)
}
