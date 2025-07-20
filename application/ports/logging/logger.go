package application

import "context"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

// Logger provides structured logging capabilities for the application.
// All methods accept a context for request tracing and structured fields
// for better log searchability and analysis.
type Logger interface {
	Debug(ctx context.Context, msg string, fields ...any)
	Info(ctx context.Context, msg string, fields ...any)
	Warn(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, fields ...any)
	// With(fields ...Field) Logger
}
