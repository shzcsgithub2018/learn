package String

import (
	"context"
	"github.com/shzgithub2018/learn/redis/dal"
)

// IdGenerator
// @Description: 实现一个递增的分布式 ID 生产器
type IdGenerator struct {
	ctx context.Context
	key string
}

func NewIdGenerator(ctx context.Context, key string) *IdGenerator {
	return &IdGenerator{
		ctx: ctx,
		key: key,
	}
}

func (h *IdGenerator) Init(startId int64) {
	dal.GetRedisCli(h.ctx).SetNX(h.ctx, h.key, startId, 0)
}

func (h *IdGenerator) GenId() int64 {
	return dal.GetRedisCli(h.ctx).Incr(h.ctx, h.key).Val()
}

func (h *IdGenerator) Reset() {
	dal.GetRedisCli(h.ctx).Del(h.ctx, h.key)
}
