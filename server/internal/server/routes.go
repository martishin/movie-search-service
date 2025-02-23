package server

import (
	"log/slog"
	"net/http"

	"github.com/martishin/movie-search-service/internal/handlers"
	"github.com/martishin/movie-search-service/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func registerRoutes(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestIDMiddleware(logger))
	r.Use(middleware.LoggingMiddleware())

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	// Root and health routes
	r.Get("/", handlers.HelloWorldHandler())

	// // API routes (protected)
	// r.With(middleware.AuthMiddleware).Get("/api/user", handlers.GetUserHandler(s.db))

	return r
}
