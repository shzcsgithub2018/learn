package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/redis/dal"
	"log"
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
		log.Println(res, " set success.")
	} else {
		log.Println(res, " set success.  ", cnt-res, " update success.")
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

// HIncrBy
// @Description: 存量、增量只能是整数
// @receiver h
// @param hashKey
// @param n
func (h *Hash) HIncrBy(hashKey string, n int64) {
	cmd := h.cli.HIncrBy(h.ctx, h.key, hashKey, n)
	if err := cmd.Err(); err != nil {
		h.t.Error(err)
	} else {
		h.t.Log(cmd.Val())
	}
}

// HIncrByFloat
// @Description: 存量、增量可以是整数，也可以是浮点数
// @receiver h
// @param hashKey
// @param n
func (h *Hash) HIncrByFloat(hashKey string, n float64) {
	cmd := h.cli.HIncrByFloat(h.ctx, h.key, hashKey, n)
	if err := cmd.Err(); err != nil {
		h.t.Error(err)
	} else {
		h.t.Log(cmd.Val())
	}
}

// HStrLen
// @Description: redis v9 不支持, 字节支持
// @receiver h
func (h *Hash) HStrLen() {

}

func (h *Hash) HExists(hashKey string) {
	log.Println("HExists", h.key, hashKey)
	if h.cli.HExists(h.ctx, h.key, hashKey).Val() {
		log.Println("Yes")
	} else {
		log.Println("NO")
	}
}

func (h *Hash) HDel(hashKeys ...string) {
	log.Print("HDel ", h.key)
	for _, hashKey := range hashKeys {
		log.Print(hashKey)
	}
	log.Println()
	cmd := h.cli.HDel(h.ctx, h.key, hashKeys...)
	if err := cmd.Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(cmd.Val(), "del success.")
	}
}

func (h *Hash) HLen() {
	log.Println("HLen", h.key)
	cmd := h.cli.HLen(h.ctx, h.key)
	if err := cmd.Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (h *Hash) HMGet(hashKeyList ...string) {
	res := h.cli.HMGet(h.ctx, h.key, hashKeyList...).Val()
	for i, _ := range res {
		log.Println(hashKeyList[i], ":", res[i])
	}
}

func (h *Hash) HKeys() {
	keys := dal.GetRedisCli(h.ctx).HKeys(h.ctx, h.key).Val()
	log.Println(keys)
}

func (h *Hash) HVals() {
	vals := dal.GetRedisCli(h.ctx).HVals(h.ctx, h.key).Val()
	log.Println(vals)
}

func (h *Hash) HGetAll() {
	allMap := dal.GetRedisCli(h.ctx).HGetAll(h.ctx, h.key).Val()
	log.Println(allMap)
}
