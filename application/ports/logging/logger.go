package application

import "context"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

type Logger interface {
	Debug(ctx context.Context, msg string, fields ...any)
	Info(ctx context.Context, msg string, fields ...any)
	Warn(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, fields ...any)
}
