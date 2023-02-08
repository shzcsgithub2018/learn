package String

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
	"time"
)

const VALUE_OF_LOCK = "locking"

type Lock struct {
	ctx        context.Context
	key        string
	expiration time.Duration
}

type OptionLock func(lock *Lock)

func NewLock(ctx context.Context, key string, opts ...OptionLock) *Lock {
	lock := &Lock{
		ctx: ctx,
		key: key,
	}
	for _, opt := range opts {
		opt(lock)
	}
	return lock
}

func WithLockExpiration(expiration time.Duration) OptionLock {
	return func(lock *Lock) {
		lock.expiration = expiration
	}
}

func (h *Lock) Acquire() bool {
	return dal.GetRedisCli(h.ctx).SetNX(h.ctx, h.key, VALUE_OF_LOCK, h.expiration).Val()
}

func (h *Lock) Release() bool {
	return dal.GetRedisCli(h.ctx).Del(h.ctx, h.key).Val() == 1
}
