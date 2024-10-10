package logging

import (
	"time"
)

// LogRequestEntry defines the structure for our request log output
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

// NewLogRequestEntry creates a new LogRequestEntry with required fields
func NewLogRequestEntry(contextID, method string, start time.Time) *LogRequestEntry {
	return &LogRequestEntry{
		ContextID: contextID,
		Package:   "handler", // Default package value
		Method:    method,
		Timestamp: start, // Default timestamp value
	}
}

// WithPackage sets a custom package name
func (b *LogRequestEntry) WithPackage(pkg string) *LogRequestEntry {
	b.Package = pkg
	return b
}

// WithMessage sets the message field
func (b *LogRequestEntry) WithMessage(message string) *LogRequestEntry {
	b.Message = message
	return b
}

// WithURL sets the URL field
func (b *LogRequestEntry) WithURL(url string) *LogRequestEntry {
	b.URL = url
	return b
}

// WithRemoteAddr sets the RemoteAddr field
func (b *LogRequestEntry) WithRemoteAddr(remoteAddr string) *LogRequestEntry {
	b.RemoteAddr = remoteAddr
	return b
}

// WithUserAgent sets the UserAgent field
func (b *LogRequestEntry) WithUserAgent(userAgent string) *LogRequestEntry {
	b.UserAgent = userAgent
	return b
}

// WithStatusCode sets the StatusCode field
func (b *LogRequestEntry) WithStatusCode(statusCode int) *LogRequestEntry {
	b.StatusCode = statusCode
	return b
}

// WithResponseTime sets the ResponseTime field
func (b *LogRequestEntry) WithResponseTime(responseTime string) *LogRequestEntry {
	b.ResponseTime = responseTime
	return b
}

// WithQueryParams sets the QueryParams field
func (b *LogRequestEntry) WithQueryParams(queryParams string) *LogRequestEntry {
	b.QueryParams = queryParams
	return b
}
