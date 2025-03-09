package middleware

import (
	"context"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/martishin/movie-search-service/internal/model/config"
)

func SessionAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user ID from the session
		userID, err := gothic.GetFromSession("user_id", r)
		if err != nil || userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Refresh the session's lifetime
		session, _ := gothic.Store.Get(r, gothic.SessionName)
		session.Options.MaxAge = 7 * 24 * 60 * 60 //nolint:mnd    // Extend by 1 week
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Failed to refresh session", http.StatusInternalServerError)
			return
		}

		// Attach user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AlloyAuthMiddleware(alloyConfig *config.ObservabilityConfig) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authUsername, authPassword, ok := r.BasicAuth()
			if !ok || authUsername != alloyConfig.AlloyUsername || authPassword != alloyConfig.AlloyPassword {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
}
