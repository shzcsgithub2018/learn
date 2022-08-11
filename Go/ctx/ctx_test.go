package ctx

import (
	"context"
	"testing"
)

var idempotentSpinKey string = "keykey"

func TestCtx(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		var retry int32
		ctx, retry = RetryCount(ctx)
		t.Log(retry)
	}
}

func RetryCount(ctx context.Context) (context.Context, int32) {
	if count, ok := ctx.Value(idempotentSpinKey).(int32); ok {
		ctx = context.WithValue(ctx, idempotentSpinKey, count+1)
		return ctx, count + 1
	}
	count := int32(1)
	ctx = context.WithValue(ctx, idempotentSpinKey, count)
	return ctx, count
}
