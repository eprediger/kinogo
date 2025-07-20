package infrastructure

import (
	application "application/ports/logging"
	"context"
	"log/slog"
	"os"
	"time"
)

type logger struct {
	logger *slog.Logger
}

func NewLogger(level application.LogLevel) application.Logger {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: translateToSlogLevel(level)})

	return &logger{
		logger: slog.New(handler),
	}
}

func (l *logger) Debug(ctx context.Context, msg string, fields ...any) {
	l.logWithLevel(ctx, slog.LevelDebug, msg, fields...)
}

func (l *logger) Error(ctx context.Context, msg string, fields ...any) {
	l.logWithLevel(ctx, slog.LevelError, msg, fields...)
}

func (l *logger) Info(ctx context.Context, msg string, fields ...any) {
	l.logWithLevel(ctx, slog.LevelInfo, msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...any) {
	l.logWithLevel(ctx, slog.LevelWarn, msg, fields...)
}

func translateToSlogLevel(level application.LogLevel) slog.Level {
	switch level {
	case application.Debug:
		return slog.LevelDebug
	case application.Info:
		return slog.LevelInfo
	case application.Warn:
		return slog.LevelWarn
	case application.Error:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *logger) logWithLevel(ctx context.Context, level slog.Level, msg string, fields ...any) {
	if !l.logger.Enabled(ctx, level) {
		return
	}

	var pc uintptr
	record := slog.NewRecord(time.Now(), level, msg, pc)
	record = record.Clone()

	// Convert fields to slog attributes
	for _, field := range fields {
		l.addFieldToRecord(&record, field)
	}

	_ = l.logger.Handler().Handle(ctx, record)
}

func (l *logger) addFieldToRecord(record *slog.Record, field any) {
	switch v := any(field).(type) {
	case slog.Attr:
		record.AddAttrs(v)
	case string:
		// Treat single strings as values with empty key
		record.AddAttrs(slog.Any("", v))
	default:
		// For other types, convert to string representation
		record.AddAttrs(slog.Any("field", v))
	}
}
