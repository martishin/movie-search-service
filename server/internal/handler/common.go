package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/martishin/movie-search-service/internal/middleware"
)

func HelloWorldHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := middleware.GetLogger(r.Context())

		resp := map[string]string{"message": "Hello World!"}

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("error handling JSON marshal", slog.Any("error", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(jsonResp)
		logger.Info("Response sent successfully", slog.String("response", string(jsonResp)))
	}
}
