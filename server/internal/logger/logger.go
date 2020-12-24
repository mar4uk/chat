package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// Logger is struct which describes logger
type Logger struct {
	Logger *zap.SugaredLogger
}

// Info is method to log string message
func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

// Infof is method to log formatted message
func (l *Logger) Infof(template string, args ...interface{}) {
	l.Logger.Infof(template, args...)
}

// Error is method to log string error
func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

// Fatal is method to log fatal error
func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

// WithFields is method to add context to log message
func (l *Logger) WithFields(args ...interface{}) *Logger {
	l.Logger = l.Logger.With(args...)
	return l
}

// NewLogger creates new instance of logger
func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Logger{sugar}, nil
}

// NewLogEntry creates logger with context, used for access logs
func (l *Logger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &Entry{Logger: l.Logger}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	entry.Logger = entry.Logger.With(
		"time", time.Now().UTC().Format(time.RFC1123),
		"http_scheme", scheme,
		"http_proto", r.Proto,
		"http_method", r.Method,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
		"uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
	)

	entry.Logger.Info("request started")
	return entry
}

// Entry describes logger
type Entry struct {
	Logger *zap.SugaredLogger
}

// Write adds context info to entry
func (l *Entry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.With(
		"status", status,
		"resp_bytes_length", bytes,
		"resp_elapsed_ms", float64(elapsed.Nanoseconds())/1000000.0,
	)
	l.Logger.Info("request complete")
}

// Panic adds context to entry in case of panic
func (l *Entry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.With(
		"stack", string(stack),
		"panic", fmt.Sprintf("%+v", v),
	)
}
