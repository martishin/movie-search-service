package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
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
		// Retrieve user ID from session
		userIDStr, err := gothic.GetFromSession("user_id", r)
		if err != nil || userIDStr == "" {
			adapter.JsonErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Convert userID from string to int
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			adapter.JsonErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
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
