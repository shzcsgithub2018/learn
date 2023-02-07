package string

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
	"strconv"
)

type Limiter struct {
	ctx context.Context
	key string
}

func NewLimiter(ctx context.Context, key string) *Limiter {
	return &Limiter{
		ctx: ctx,
		key: key,
	}
}

func (h *Limiter) SetMaxExecuteTimes(maxExecuteTimes int64) {
	dal.GetRedisCli(h.ctx).Set(h.ctx, h.key, maxExecuteTimes, 0)
}

func (h *Limiter) StillValidToExecute() bool {
	val := dal.GetRedisCli(h.ctx).Decr(h.ctx, h.key).Val()
	return val >= 0
}

func (h *Limiter) RemainingExecuteTimes() int64 {
	valStr := dal.GetRedisCli(h.ctx).Get(h.ctx, h.key).Val()
	val, _ := strconv.ParseInt(valStr, 10, 64)
	return val
}
