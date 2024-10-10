package logging

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logger interface {
	LogRequest(entry *LogRequestEntry)
	LogInternal(ctx context.Context, packg, method, message string, args ...interface{})
	LogError(ctx context.Context, packg, method string, err error)
}

type logger struct {
	timestampFormat string
}

var instance *logger
var once sync.Once

// NewLogger creates a new logger instance (singleton)
func NewLogger() Logger {
	once.Do(func() {
		instance = &logger{
			timestampFormat: time.RFC3339Nano,
		}
	})
	return instance
}

// GetLogger gets the logger instance
func GetLogger() Logger {
	return instance
}

// LogRequest logs messages with a standard format for requests
func (l *logger) LogRequest(entry *LogRequestEntry) {
	var sb strings.Builder

	if entry.Package == "" {
		entry.Package = string(defaultPackage)
	}

	// Start building the log message
	sb.WriteString("context=")
	sb.WriteString(entry.ContextID)
	sb.WriteString(" package=")
	sb.WriteString(entry.Package)
	sb.WriteString(" method=")
	sb.WriteString(entry.Method)
	sb.WriteString(" timestamp=")
	sb.WriteString(entry.Timestamp.Format(instance.timestampFormat))
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

	if entry.QueryParams != "" {
		sb.WriteString(" query_params=")
		sb.WriteString(entry.QueryParams)
	}

	// Log the final string
	log.Println(sb.String())
}

// LogInternal logs messages with a standard format for the internal packages
func (l *logger) LogInternal(ctx context.Context, packg, method, message string, args ...interface{}) {
	var sb strings.Builder

	// Build the log message
	sb.WriteString("context=")
	sb.WriteString(fmt.Sprintf("%v", ctx.Value(ContextIDKey)))
	sb.WriteString(" package=")
	sb.WriteString(packg)
	sb.WriteString(" method=")
	sb.WriteString(method)
	sb.WriteString(" timestamp=")
	sb.WriteString(time.Now().Format(instance.timestampFormat))
	sb.WriteString(" ")
	sb.WriteString(formatMessage(message, args...))

	// Log the final message
	log.Println(sb.String())
}

// LogError logs messages with a standard format for the errors
func (l *logger) LogError(ctx context.Context, packg, method string, err error) {
	var sb strings.Builder

	// Build the log message
	sb.WriteString("context=")
	sb.WriteString(fmt.Sprintf("%v", ctx.Value(ContextIDKey)))
	sb.WriteString(" package=")
	sb.WriteString(packg)
	sb.WriteString(" method=")
	sb.WriteString(method)
	sb.WriteString(" timestamp=")
	sb.WriteString(time.Now().Format(instance.timestampFormat))
	sb.WriteString(" ")
	sb.WriteString(err.Error()) // Convert the error to string

	// Log the final message
	log.Println(sb.String())
}

// Helper function to format message
func formatMessage(message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}
