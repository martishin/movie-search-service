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

func RegisterRoutes(
	logger *slog.Logger,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	movieHandler *handler.MovieHandler,
) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestIDMiddleware(logger))
	r.Use(middleware.LoggingMiddleware())

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://ms.martishin.com"},
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
		api.Post("/logout", authHandler.LogoutHandler())
		api.Post("/signup", authHandler.SignUpHandler())
		api.Post("/login", authHandler.LoginHandler())
	})

	// API routes (protected)
	r.Route("/api", func(api chi.Router) {
		api.With(middleware.AuthMiddleware).Get("/users/me", userHandler.GetUserHandler())

		// Movie endpoints
		api.Get("/movies", movieHandler.ListMoviesHandler())
		api.Get("/movies/{id}", movieHandler.GetMovieHandler())
		api.Post("/movies", movieHandler.CreateMovieHandler())
		api.Get("/genres", movieHandler.ListGenresHandler())

		// Liking Movies
		api.Route("/movies/likes", func(likeRouter chi.Router) {
			likeRouter.Use(middleware.AuthMiddleware)

			likeRouter.Get("/", userHandler.GetLikedMoviesHandler())
			likeRouter.Post("/{movie_id}", userHandler.AddLikeHandler())
			likeRouter.Delete("/{movie_id}", userHandler.RemoveLikeHandler())
		})

		// Movies with likes
		api.Route("/movies-with-likes", func(moviesWithLikesRouter chi.Router) {
			moviesWithLikesRouter.Use(middleware.AuthMiddleware)

			moviesWithLikesRouter.Get("/", movieHandler.ListMoviesWithGenresAndLikesHandler())
			moviesWithLikesRouter.Get("/{movie_id}", movieHandler.GetMovieHandlerWithLike())
		})

		// Admin endpoints
		api.Route("/admin", func(admin chi.Router) {
			admin.Use(middleware.AuthMiddleware)

			admin.Post("/movies", movieHandler.CreateMovieHandler())
			admin.Put("/movies/{id}", movieHandler.UpdateMovieHandler())
			admin.Delete("/movies/{id}", movieHandler.DeleteMovieHandler())
		})
	})

	return r
}
