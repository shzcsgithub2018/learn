package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

type Hash struct {
	ctx context.Context
	t   *testing.T
	key string
	cli *redis.Client
}

func NewHash(ctx context.Context, t *testing.T, key string, cli *redis.Client) *Hash {
	cli.Del(ctx, key)
	return &Hash{
		ctx: ctx,
		t:   t,
		key: key,
		cli: cli,
	}
}

func (h *Hash) HSet(values ...interface{}) {
	cnt := int64(len(values) / 2)
	res := h.cli.HSet(h.ctx, h.key, values).Val()
	if cnt == res {
		h.t.Log(res, " set success.")
	} else {
		h.t.Log(res, " set success.  ", cnt-res, " update success.")
	}
}

func (h *Hash) HSetNX(hashKey, hashVal string) {
	success := h.cli.HSetNX(h.ctx, h.key, hashKey, hashVal).Val()
	if success {
		h.t.Log("success.")
	} else {
		h.t.Log("fail.")
	}
}

func (h *Hash) HGet(hashKey string) {
	cmd := h.cli.HGet(h.ctx, h.key, hashKey)
	if err := cmd.Err(); err != nil {
		h.t.Error(err)
	} else {
		h.t.Log(cmd.Val())
	}
}
