package route

import (
	"log/slog"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/martishin/movie-search-service/internal/handler"
	"github.com/martishin/movie-search-service/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RegisterRoutes(logger *slog.Logger, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) http.Handler {
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
	r.Get("/", handler.HelloWorldHandler())

	// Authentication routes
	r.Route("/auth", func(api chi.Router) {
		api.Get("/start", gothic.BeginAuthHandler)
		api.Get("/callback", authHandler.GoogleCallbackHandler())
		api.Get("/logout", authHandler.LogoutHandler())
		api.Post("/signup", authHandler.SignUpHandler())
		api.Post("/login", authHandler.LoginHandler())
	})

	// API routes (protected)
	r.Route("/api", func(api chi.Router) {
		api.With(middleware.AuthMiddleware).Get("/users/me", userHandler.GetUserHandler())
	})

	return r
}
