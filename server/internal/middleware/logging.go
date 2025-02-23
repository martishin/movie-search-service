package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs request details (with body) and response (without body).
func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Get request-specific logger
			logger := GetLogger(r.Context())

			// Read request body
			var requestBody []byte
			if r.Body != nil {
				requestBody, _ = io.ReadAll(r.Body)
				r.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // Restore body
			}

			// Log incoming request (including request body)
			logger.Info("Request received",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("request_body", string(requestBody)),
			)

			// Wrap ResponseWriter to capture status code
			responseWrapper := &responseWriterWrapper{
				ResponseWriter: w,
				statusCode:     http.StatusOK, // Default to 200
			}

			// Process request
			next.ServeHTTP(responseWrapper, r)

			durationMs := time.Since(start).Milliseconds()

			// Log response (without response body)
			logger.Info("Request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", responseWrapper.statusCode),
				slog.Int64("duration_ms", durationMs),
			)
		})
	}
}
