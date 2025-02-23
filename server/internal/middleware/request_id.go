package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

const requestIDKey = "requestID"

// RequestIDMiddleware generates a request ID and injects it into the request context and logger.
func RequestIDMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Generate a unique request ID
			reqID := uuid.New().String()

			// Store it in the request context
			ctx := context.WithValue(r.Context(), requestIDKey, reqID)

			// Create a new logger with request ID pre-attached
			requestLogger := logger.With(slog.String("request_id", reqID))

			// Store the new logger in context
			ctx = context.WithValue(ctx, "logger", requestLogger)

			// Add Request ID to response headers
			w.Header().Set("X-Request-ID", reqID)

			// Pass the updated context to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value("logger").(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return "unknown"
}
