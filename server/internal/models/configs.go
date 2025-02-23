package models

import (
	"time"
)

type ServerConfig struct {
	Port         int
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type OAuthConfig struct {
	ClientID      string
	ClientSecret  string
	CallbackURL   string
	SessionSecret string
	Domain        string
	IsProduction  bool
}
