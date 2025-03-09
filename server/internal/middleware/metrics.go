package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
)

// Define Prometheus Metrics
var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response duration for requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status", "env"},
	)

	requestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route", "status", "env"},
	)

	inFlightRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of in-flight requests",
		},
		[]string{"method", "route", "env"},
	)

	errorTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_errors_total",
			Help: "Total number of HTTP requests resulting in errors",
		},
		[]string{"method", "route", "status", "env"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration, requestTotal, inFlightRequests, errorTotal)
}

func MetricsMiddleware() func(http.Handler) http.Handler {
	env := os.Getenv("ENV")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Track in-flight requests
			inFlightRequests.WithLabelValues(r.Method, r.URL.Path, env).Inc()
			defer inFlightRequests.WithLabelValues(r.Method, r.URL.Path, env).Dec()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			duration := time.Since(start).Seconds()
			status := strconv.Itoa(ww.Status())

			// Record metrics
			requestDuration.WithLabelValues(r.Method, r.URL.Path, status, env).Observe(duration)
			requestTotal.WithLabelValues(r.Method, r.URL.Path, status, env).Inc()

			// Track errors (4xx, 5xx)
			if ww.Status() >= 400 {
				errorTotal.WithLabelValues(r.Method, r.URL.Path, status, env).Inc()
			}
		})
	}
}
