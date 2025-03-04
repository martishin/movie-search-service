package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/martishin/movie-search-service/internal/adapter"
	"github.com/martishin/movie-search-service/internal/middleware"
	"github.com/martishin/movie-search-service/internal/model/domain"
	"github.com/martishin/movie-search-service/internal/service"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (h *MovieHandler) CreateMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		var request domain.Movie
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("Invalid request payload", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Invalid request", http.StatusBadRequest)
			return
		}

		movie, err := h.movieService.CreateMovie(r.Context(), request)
		if err != nil {
			logger.Error("Failed to create movie", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not create movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movie)
	}
}

func (h *MovieHandler) GetMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		idStr := r.PathValue("id")
		movieID, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("Invalid movie ID", slog.String("id", idStr))
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		movie, err := h.movieService.GetMovieByIDWithGenres(r.Context(), movieID)
		if err != nil {
			logger.Error("Movie not found", slog.String("id", idStr))
			adapter.JsonErrorResponse(w, "Movie not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie)
	}
}

func (h *MovieHandler) GetMovieHandlerWithLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		movieID, err := strconv.Atoi(r.PathValue("movie_id"))
		if err != nil {
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		movie, err := h.movieService.GetMovieByIDWithGenresAndLike(r.Context(), movieID, userID)
		if err != nil {
			adapter.JsonErrorResponse(w, "Movie not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie)
	}
}

func (h *MovieHandler) ListMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		movies, err := h.movieService.ListMoviesWithGenres(r.Context())
		if err != nil {
			logger.Error("Failed to fetch movies", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not fetch movies", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	}
}

func (h *MovieHandler) UpdateMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		idStr := r.PathValue("id")
		movieID, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("Invalid movie ID", slog.String("id", idStr))
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		var request domain.Movie
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("Invalid request payload", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Invalid request", http.StatusBadRequest)
			return
		}

		request.ID = movieID // Ensure the correct ID is set

		err = h.movieService.UpdateMovie(r.Context(), request)
		if err != nil {
			logger.Error("Failed to update movie", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not update movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Movie updated successfully"})
	}
}

func (h *MovieHandler) DeleteMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		idStr := r.PathValue("id")
		movieID, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("Invalid movie ID", slog.String("id", idStr))
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		err = h.movieService.DeleteMovie(r.Context(), movieID)
		if err != nil {
			logger.Error("Failed to delete movie", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not delete movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted successfully"})
	}
}

func (h *MovieHandler) ListGenresHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		genres, err := h.movieService.ListGenres(r.Context())
		if err != nil {
			logger.Error("Failed to fetch genres", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not fetch genres", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(genres)
	}
}

func (h *MovieHandler) ListMoviesWithGenresAndLikesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		movies, err := h.movieService.ListMoviesWithGenresAndLikes(r.Context(), userID)
		if err != nil {
			logger.Error("Failed to fetch movies with genres and likes", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not fetch movies", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	}
}
