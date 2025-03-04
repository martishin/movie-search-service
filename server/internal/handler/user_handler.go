package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/martishin/movie-search-service/internal/adapter"
	"github.com/martishin/movie-search-service/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Fetch user from service
		user, err := h.userService.GetUserByID(r.Context(), userID)
		if err != nil {
			adapter.JsonErrorResponse(w, "User not found", http.StatusNotFound)
			return
		}

		// Send response
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserHandler) AddLikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		movieID, err := strconv.Atoi(chi.URLParam(r, "movie_id"))
		if err != nil {
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		if err := h.userService.LikeMovie(r.Context(), userID, movieID); err != nil {
			adapter.JsonErrorResponse(w, "Could not like movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *UserHandler) RemoveLikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		movieID, err := strconv.Atoi(chi.URLParam(r, "movie_id"))
		if err != nil {
			adapter.JsonErrorResponse(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}

		if err := h.userService.UnlikeMovie(r.Context(), userID, movieID); err != nil {
			adapter.JsonErrorResponse(w, "Could not unlike movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *UserHandler) GetLikedMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := adapter.GetUserIDFromSession(r)
		if err != nil {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		movies, err := h.userService.GetLikedMovies(r.Context(), userID)
		if err != nil {
			adapter.JsonErrorResponse(w, "Could not fetch liked movies", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(movies)
	}
}
