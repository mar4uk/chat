package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Logger{sugar}, nil
}

func (l *Logger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &LoggerEntry{Logger: l.Logger}

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

type LoggerEntry struct {
	Logger *zap.SugaredLogger
}

func (l *LoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.With(
		"status", status,
		"resp_bytes_length", bytes,
		"resp_elapsed_ms", float64(elapsed.Nanoseconds())/1000000.0,
	)
	l.Logger.Info("request complete")
}

func (l *LoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.With(
		"stack", string(stack),
		"panic", fmt.Sprintf("%+v", v),
	)
}
