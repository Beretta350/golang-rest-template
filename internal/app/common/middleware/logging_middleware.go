package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/Beretta350/golang-rest-template/pkg/logging"
	"github.com/google/uuid"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

// LoggingMiddleware logs incoming requests and outgoing responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a unique request ID
		requestID := uuid.New().String()

		start := time.Now()

		// Wrap the ResponseWriter to capture the status code
		crw := &CustomResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}

		ctx := context.WithValue(r.Context(), logging.ContextIDKey, requestID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(crw, r)

		// Calculate response time
		duration := time.Since(start)

		// Build log entry
		logEntry := logging.NewLogRequestEntry(requestID, r.Method, start).
			WithURL(r.URL.Path).
			WithRemoteAddr(r.RemoteAddr).
			WithUserAgent(r.UserAgent()).
			WithStatusCode(crw.StatusCode).
			WithResponseTime(duration.String()).
			WithQueryParams(r.URL.RawQuery)

		// Log in structured JSON format
		logging.GetLogger().LogRequest(logEntry)
	})
}
