package middleware

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is a struct for standard error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandlerMiddleware is the middleware that handles errors
func ErrorHandlerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler and check if an error occurred
		defer func() {
			if err := recover(); err != nil {
				// Log the error (optional)
				// You can log the error here if needed

				// Send a 500 Internal Server Error response
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal Server Error"})
			}
		}()

		next.ServeHTTP(w, r)
	}
}
