package logging

import (
	"time"
)

// LogEntry defines the structure for our request log output
type LogRequestEntry struct {
	ContextID    string    `json:"context"`
	Package      string    `json:"package"`
	Method       string    `json:"method"`
	Message      string    `json:"message,omitempty"`
	URL          string    `json:"url,omitempty"`
	RemoteAddr   string    `json:"remote_addr,omitempty"`
	UserAgent    string    `json:"user_agent,omitempty"`
	StatusCode   int       `json:"status_code,omitempty"`
	ResponseTime string    `json:"response_time,omitempty"`
	QueryParams  string    `json:"query_params,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
}
