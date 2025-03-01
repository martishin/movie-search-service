package config

import "time"

type ServerConfig struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
