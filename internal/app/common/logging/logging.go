package logging

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Beretta350/golang-rest-template/internal/app/common/constants"
)

// LogEntry defines the structure for our log output
type LogEntry struct {
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

// LogRequest: logs messages with a standard format for the handler package
func LogRequest(entry LogEntry) {
	var sb strings.Builder

	// Start building the log message
	sb.WriteString("context=")
	sb.WriteString(entry.ContextID)
	sb.WriteString(" package=handler")
	sb.WriteString(" method=")
	sb.WriteString(entry.Method)
	sb.WriteString(" timestamp=")
	sb.WriteString(entry.Timestamp.Format(time.RFC3339))
	sb.WriteString(" url=")
	sb.WriteString(entry.URL)
	sb.WriteString(" remote_addr=")
	sb.WriteString(entry.RemoteAddr)
	sb.WriteString(" user_agent=")
	sb.WriteString(entry.UserAgent)
	sb.WriteString(" status_code=")
	sb.WriteString(strconv.Itoa(entry.StatusCode))
	sb.WriteString(" response_time=")
	sb.WriteString(entry.ResponseTime)
	sb.WriteString(" query_params=")
	sb.WriteString(entry.QueryParams)

	// Log the final string
	log.Println(sb.String())
}

// LogService: logs messages with a standard format for the service package
func LogService(ctx context.Context, method, message string, args ...interface{}) {
	var sb strings.Builder

	// Build the log message
	sb.WriteString("context=")
	sb.WriteString(fmt.Sprintf("%v", ctx.Value(constants.RequestIDKey)))
	sb.WriteString(" package=service method=")
	sb.WriteString(method)
	sb.WriteString(" timestamp=")
	sb.WriteString(time.Now().Format(time.RFC3339))
	sb.WriteString(" ")
	sb.WriteString(formatMessage(message, args...))

	// Log the final message
	log.Println(sb.String())
}

// LogService: logs messages with a standard format for the service package
func LogError(ctx context.Context, method string, err error) {
	var sb strings.Builder

	// Build the log message
	sb.WriteString("context=")
	sb.WriteString(fmt.Sprintf("%v", ctx.Value(constants.RequestIDKey)))
	sb.WriteString(" package=service method=")
	sb.WriteString(method)
	sb.WriteString(" timestamp=")
	sb.WriteString(time.Now().Format(time.RFC3339))
	sb.WriteString(" ")
	sb.WriteString(err.Error()) // Convert the error to string

	// Log the final message
	log.Println(sb.String())
}

// Helper function to format the message
func formatMessage(message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}
