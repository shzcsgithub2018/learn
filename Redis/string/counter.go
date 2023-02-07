package string

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
	"strconv"
)

type Counter struct {
	ctx  context.Context
	key  string
	step int64
}

func NewCounter(ctx context.Context, key string) *Counter {
	return &Counter{
		ctx:  ctx,
		key:  key,
		step: 1,
	}
}

type Option func(c *Counter)

func WithModifyNum(num int64) Option {
	return func(c *Counter) {
		c.step = num
	}
}

// Incr
// @Description: 默认增量 1
// @receiver c
// @param opts
// @return int64
func (c *Counter) Incr(opts ...Option) int64 {
	for _, opt := range opts {
		opt(c)
	}
	if c.step == 1 {
		return dal.GetRedisCli(c.ctx).Incr(c.ctx, c.key).Val()
	} else {
		step := c.step
		c.step = 1
		return dal.GetRedisCli(c.ctx).IncrBy(c.ctx, c.key, step).Val()
	}
}

// Decr
// @Description: 默认增量 1
// @receiver c
// @param opts
// @return int64
func (c *Counter) Decr(opts ...Option) int64 {
	for _, opt := range opts {
		opt(c)
	}
	if c.step == 1 {
		return dal.GetRedisCli(c.ctx).Decr(c.ctx, c.key).Val()
	} else {
		step := c.step
		c.step = 1
		return dal.GetRedisCli(c.ctx).DecrBy(c.ctx, c.key, step).Val()
	}
}

// Get
// @Description:
// @receiver c
// @return int64
func (c *Counter) Get() int64 {
	valStr := dal.GetRedisCli(c.ctx).Get(c.ctx, c.key).Val()
	val, _ := strconv.ParseInt(valStr, 10, 64)
	return val
}

// Reset
// @Description:
// @receiver c
// @return int64
func (c *Counter) Reset() int64 {
	cmd := dal.GetRedisCli(c.ctx).GetSet(c.ctx, c.key, 0)
	if err := cmd.Err(); err != nil {
		return 0
	}
	val, _ := strconv.ParseInt(cmd.Val(), 10, 64)
	return val
}
