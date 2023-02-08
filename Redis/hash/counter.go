package hash

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
	"log"
	"strconv"
)

// Counter
// @Description: 允许用户将多个关联的计数器存储到同一个散列中实行集中管理。并可以对指定的计数器执行加减法，不会影响其他计数器。
type Counter struct {
	ctx  context.Context
	key  string
	step int64
}

func NewCounter(ctx context.Context, key string) *Counter {
	dal.GetRedisCli(ctx).Del(ctx, key)
	log.Println("del ", key)
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
func (c *Counter) Incr(hashKey string, opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	step := c.step
	c.step = 1
	log.Println("HIncrBy "+c.key+" "+hashKey+" ", step)
	log.Println(dal.GetRedisCli(c.ctx).HIncrBy(c.ctx, c.key, hashKey, step).Val())
}

// Decr
// @Description: 默认增量 1
// @receiver c
// @param opts
// @return int64
func (c *Counter) Decr(hashKey string, opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	step := c.step
	c.step = 1
	log.Println("HIncrBy "+c.key+" "+hashKey+" ", -step)
	log.Println(dal.GetRedisCli(c.ctx).HIncrBy(c.ctx, c.key, hashKey, -step).Val())
}

// Get
// @Description:
// @receiver c
// @return int64
func (c *Counter) Get(hashKey string) int64 {
	valStr := dal.GetRedisCli(c.ctx).HGet(c.ctx, c.key, hashKey).Val()
	val, _ := strconv.ParseInt(valStr, 10, 64)
	log.Println("HGet ", c.key, hashKey)
	log.Println(val)
	return val
}

// Reset
// @Description:
// @receiver c
func (c *Counter) Reset(hashKey string) {
	c.Get(hashKey)
	cmd := dal.GetRedisCli(c.ctx).HSet(c.ctx, c.key, hashKey, 0)
	if err := cmd.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("HSet", c.key, hashKey, 0)
}
