package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
	"github.com/martishin/movie-search-service/internal/adapter"
	"github.com/martishin/movie-search-service/internal/middleware"
	"github.com/martishin/movie-search-service/internal/model/config"
	"github.com/martishin/movie-search-service/internal/service"
	"golang.org/x/crypto/bcrypt"
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
			adapter.JsonErrorResponse(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		// Save user to database
		ctx := r.Context()
		user, err := h.userService.FindOrCreateUser(ctx, authUser.FirstName, authUser.LastName, authUser.Email, authUser.AvatarURL, "")
		if err != nil {
			logger.Error("Failed to find or create user", slog.Any("error", err), slog.String("email", authUser.Email))
			adapter.JsonErrorResponse(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Convert user ID to string for session storage
		userIDStr := strconv.Itoa(user.ID)

		// Store user ID in session
		err = gothic.StoreInSession("user_id", userIDStr, r, w)
		if err != nil {
			logger.Error("Failed to store user ID in session", slog.Any("error", err), slog.String("user_id", userIDStr))
			adapter.JsonErrorResponse(w, "Failed to save session", http.StatusInternalServerError)
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
			adapter.JsonErrorResponse(w, "Failed to logout", http.StatusInternalServerError)
			return
		}

		// Redirect user
		http.Redirect(w, r, h.oauthConfig.RedirectURL, http.StatusTemporaryRedirect)
	}
}

// SignUpHandler handles user registration
func (h *AuthHandler) SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		// Parse JSON request
		var request struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("Invalid request payload", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Invalid request", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		// Check if user already exists
		_, err := h.userService.GetUserByEmail(ctx, request.Email)
		if err == nil {
			adapter.JsonErrorResponse(w, "User already exists", http.StatusConflict)
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			logger.Error("Failed to hash password", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create new user
		user, err := h.userService.CreateUser(ctx, request.FirstName, request.LastName, request.Email, "", string(hashedPassword))
		if err != nil {
			logger.Error("Failed to create user", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Could not create user", http.StatusInternalServerError)
			return
		}

		// Store user ID in session
		userIDStr := strconv.Itoa(user.ID)
		err = gothic.StoreInSession("user_id", userIDStr, r, w)
		if err != nil {
			logger.Error("Failed to store session", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Failed to save session", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}

// LoginHandler authenticates users
func (h *AuthHandler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("Invalid request payload", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Validate user
		ctx := r.Context()
		userID, password, err := h.userService.GetUserIDAndPasswordByEmail(ctx, request.Email)
		if err != nil {
			logger.Error("User not found", slog.String("email", request.Email))
			adapter.JsonErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Check if user is an OAuth user (i.e., no password set)
		if password == "" {
			logger.Warn("Attempted password login for OAuth user", slog.String("email", request.Email))
			adapter.JsonErrorResponse(w, "This account uses Google authentication. Please log in with Google.", http.StatusUnauthorized)
			return
		}

		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(request.Password))
		if err != nil {
			logger.Error("Invalid password attempt", slog.String("email", request.Email))
			adapter.JsonErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Store user ID in session
		userIDStr := strconv.Itoa(userID)
		err = gothic.StoreInSession("user_id", userIDStr, r, w)
		if err != nil {
			logger.Error("Failed to store session", slog.Any("error", err))
			adapter.JsonErrorResponse(w, "Failed to save session", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	}
}
