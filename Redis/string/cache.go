package String

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	ctx    context.Context
	client *redis.Client
}

func NewCache(ctx context.Context, client *redis.Client) *Cache {
	return &Cache{
		ctx:    ctx,
		client: client,
	}
}

func (h *Cache) Set(key, val string) {
	h.client.Set(h.ctx, key, val, 0)
}

func (h *Cache) Get(key string) string {
	return h.client.Get(h.ctx, key).Val()
}

func (h *Cache) Update(key, newVal string) string {
	return h.client.GetSet(h.ctx, key, newVal).Val()
}
