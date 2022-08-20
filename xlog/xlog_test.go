package xlog

import (
	"context"
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		ctx  context.Context
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				args: []any{
					"test",
					123,
				},
			},
		},
	}

	New(NewZapLogger())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.ctx, tt.args.args...)
		})
	}
}
