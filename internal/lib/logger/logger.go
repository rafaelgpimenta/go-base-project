package logger

import (
	"context"
	"log/slog"
	"os"
	"resource-management/internal/lib/tracinghook"
	"runtime/debug"
)

var Logger *slog.Logger
var TraceKey = tracinghook.TraceKey

func init() {
	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	handler := &tracinghook.TracingHandler{Handler: baseHandler}

	Logger = slog.New(handler)

	slog.SetDefault(Logger)
}

func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...any) {
	Logger.InfoContext(ctx, msg, args...)
}

func Error(msg string, err error, args ...any) {
	if err != nil {
		args = append(args, slog.Any("error", err))
		args = append(args, slog.Any("stack", string(debug.Stack())))
	}
	Logger.Error(msg, args...)
}

func ErrorCtx(ctx context.Context, msg string, err error, args ...any) {
	if err != nil {
		args = append(args, slog.Any("error", err))
		args = append(args, slog.Any("stack", string(debug.Stack())))
	}
	Logger.ErrorContext(ctx, msg, args...)
}

func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}

func DebugCtx(ctx context.Context, msg string, args ...any) {
	Logger.DebugContext(ctx, msg, args...)
}

func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func WarnCtx(ctx context.Context, msg string, args ...any) {
	Logger.WarnContext(ctx, msg, args...)
}
