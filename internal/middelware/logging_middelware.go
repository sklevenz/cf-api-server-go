package middleware

import (
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/logger"
)

// LoggingMiddleware logs all incoming HTTP requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
