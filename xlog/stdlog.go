package xlog

import (
	"context"
	"log"
)

type StdLogger struct{}

func NewStdLogger() Logger {
	return &StdLogger{}
}

func (l *StdLogger) Info(ctx context.Context, args ...any) {
	infos := make([]any, 0)
	infos = append(infos, "INFO")
	infos = append(infos, args...)
	log.Println(infos...)
}

func (l *StdLogger) Debug(ctx context.Context, args ...any) {
	infos := make([]any, 0)
	infos = append(infos, "INFO")
	infos = append(infos, args...)
	log.Println(infos...)
}

func (l *StdLogger) Error(ctx context.Context, args ...any) {
	infos := make([]any, 0)
	infos = append(infos, "INFO")
	infos = append(infos, args...)
	log.Println(infos...)
}

func (l *StdLogger) Panic(ctx context.Context, args ...any) {
	infos := make([]any, 0)
	infos = append(infos, "INFO")
	infos = append(infos, args...)
	log.Panic(infos...)
}
