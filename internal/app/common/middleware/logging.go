package middleware

import (
	"log"
	"time"
)

// LogEntry defines the structure for our log output
type LogEntry struct {
	Timestamp    time.Time `json:"timestamp"`
	Method       string    `json:"method"`
	URL          string    `json:"url"`
	RemoteAddr   string    `json:"remote_addr"`
	UserAgent    string    `json:"user_agent"`
	StatusCode   int       `json:"status_code"`
	ResponseTime string    `json:"response_time"`
	QueryParams  string    `json:"query_params,omitempty"`
}

// logToJSON logs the request/response details in a structured JSON format
func logRequest(entry LogEntry) {
	// Print the JSON log (this can be replaced with any logging framework)
	log.Printf("package=handler method=%s url=%s remote_addr=%s user_agent=%s status_code=%d response_time=%s query_params=%s timestamp=%s",
		entry.Method,
		entry.URL,
		entry.RemoteAddr,
		entry.UserAgent,
		entry.StatusCode,
		entry.ResponseTime,
		entry.QueryParams,
		entry.Timestamp.Format(time.RFC3339),
	)
}
