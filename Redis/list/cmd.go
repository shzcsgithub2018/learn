package list

import (
	"github.com/shzgithub2018/learn/Redis/dal"
	"log"
)

type List struct {
	key string
	cli *redis.Client
}

func NewList(key string) *List {
	dal.GetRedisCli(ctx).Del(ctx, key)
	return &List{
		key: key,
		cli: dal.GetRedisCli(ctx),
	}
}

func (h *List) LPush(vals ...string) {
	log.Println("LPush ", h.key, " ", vals)
	cnt := h.cli.LPush(ctx, h.key, vals).Val()
	log.Println("including ", cnt, " elements")
}

func (h *List) RPush(vals ...string) {
	log.Println("RPush ", h.key, " ", vals)
	cnt := h.cli.RPush(ctx, h.key, vals).Val()
	log.Println("including ", cnt, " elements")
}

func (h *List) LLen() int64 {
	log.Println("LLen ", h.key)
	lLen := h.cli.LLen(ctx, h.key).Val()
	log.Println(lLen)
	return lLen
}

func (h *List) LRange(start, end int64) {
	log.Println("LRange ", h.key, " ", start, " ", end)
	list := h.cli.LRange(ctx, h.key, start, end).Val()
	log.Println(list)
}

func (h *List) LIndex(idx int64) {
	log.Println("LRange ", h.key, " ", idx)
	element := h.cli.LIndex(ctx, h.key, idx).Val()
	log.Println(element)
}
