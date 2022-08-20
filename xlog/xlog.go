package xlog

import (
	"context"
	"sync"
)

var (
	globalLogger Logger
	once         sync.Once
)

type Logger interface {
	Info(ctx context.Context, args ...any)
	Debug(ctx context.Context, args ...any)
	Error(ctx context.Context, args ...any)
	Panic(ctx context.Context, args ...any)
}

func New(logger Logger) {
	once.Do(func() {
		globalLogger = logger
	})
}

func Info(ctx context.Context, args ...any) {
	globalLogger.Info(ctx, args...)
}

func Debug(ctx context.Context, args ...any) {
	globalLogger.Debug(ctx, args...)
}

func Error(ctx context.Context, args ...any) {
	globalLogger.Debug(ctx, args...)
}

func Panic(ctx context.Context, args ...any) {
	globalLogger.Panic(ctx, args...)
}
