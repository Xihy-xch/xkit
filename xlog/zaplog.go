package xlog

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

var (
	zapOnce sync.Once
)

type ZapLogger struct {
}

func NewZapLogger() Logger {
	zapOnce.Do(func() {
		l, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(l)
	})

	return &ZapLogger{}
}

func (l *ZapLogger) Info(ctx context.Context, args ...any) {
	zap.S().Infow("msg", args...)
}

func (l *ZapLogger) Debug(ctx context.Context, args ...any) {
	zap.S().Debugw("msg", args...)
}

func (l *ZapLogger) Error(ctx context.Context, args ...any) {
	zap.S().Errorw("msg", args...)
}

func (l *ZapLogger) Panic(ctx context.Context, args ...any) {
	zap.S().Panicw("msg", args...)
}
