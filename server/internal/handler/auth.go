package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
	"github.com/martishin/movie-search-service/internal/middleware"
	"github.com/martishin/movie-search-service/internal/model/config"
	"github.com/martishin/movie-search-service/internal/service"
)

type AuthHandler struct {
	userService *service.UserService
	oauthConfig *config.OAuthConfig
}

func NewAuthHandler(userService *service.UserService, oauthConfig *config.OAuthConfig) *AuthHandler {
	return &AuthHandler{userService: userService, oauthConfig: oauthConfig}
}

func (h *AuthHandler) GoogleCallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		// Complete authentication process
		authUser, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			logger.Error("Failed to auth user", slog.Any("error", err))
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		// Save user to database
		ctx := r.Context()
		user, err := h.userService.FindOrCreateUser(ctx, authUser.Name, authUser.Email, authUser.AvatarURL)
		if err != nil {
			logger.Error("Failed to find or create user", slog.Any("error", err), slog.String("email", authUser.Email))
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Convert user ID to string for session storage
		userIDStr := strconv.Itoa(user.ID)

		// Store user ID in session
		err = gothic.StoreInSession("user_id", userIDStr, r, w)
		if err != nil {
			logger.Error("Failed to store user ID in session", slog.Any("error", err), slog.String("user_id", userIDStr))
			http.Error(w, "Failed to save session", http.StatusInternalServerError)
			return
		}

		// Redirect user
		http.Redirect(w, r, h.oauthConfig.RedirectURL, http.StatusFound)
	}
}

// LogoutHandler logs users out and logs events properly
func (h *AuthHandler) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		// Clear session
		err := gothic.Logout(w, r)
		if err != nil {
			logger.Error("Failed to logout user", slog.Any("error", err))
			http.Error(w, "Failed to logout", http.StatusInternalServerError)
			return
		}

		// Redirect user
		http.Redirect(w, r, h.oauthConfig.RedirectURL, http.StatusTemporaryRedirect)
	}
}
