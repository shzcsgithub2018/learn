package list

import (
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/redis/dal"
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

func (h *List) LPushX(val string) {
	log.Println("LPushX ", h.key, " ", val)
	cmd := h.cli.LPushX(ctx, h.key, val)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("including ", cmd.Val(), " elements")
	}

}

func (h *List) RPushX(val string) {
	log.Println("RPushX ", h.key, " ", val)
	cmd := h.cli.RPushX(ctx, h.key, val)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("including ", cmd.Val(), " elements")
	}
}

func (h *List) LPop() {
	log.Println("LPop ", h.key)
	cmd := h.cli.LPop(ctx, h.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (h *List) RPop() {
	log.Println("RPop ", h.key)
	cmd := h.cli.RPop(ctx, h.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (h *List) RPopLPush(key string) {
	log.Println("RPopLPush ", h.key, " ", key)
	h.cli.RPopLPush(ctx, h.key, key)
}

func (h *List) LSet(idx int64, element string) {
	log.Println("LSet ", h.key, " ", idx, " ", element)
	cmd := h.cli.LSet(ctx, h.key, idx, element)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (h *List) LInsert(op, pivot, value string) {
	log.Println("LInsert ", h.key, " ", op, " ", pivot, " ", value)
	cmd := h.cli.LInsert(ctx, h.key, op, pivot, value)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("including ", cmd.Val(), " elements")
	}
}

// LTrim
// @Description: 闭区间
// @receiver h
// @param start
// @param end
func (h *List) LTrim(start, end int64) {
	log.Println("LTrim ", h.key, " ", start, " ", end)
	cmd := h.cli.LTrim(ctx, h.key, start, end)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (h *List) LRem(cnt int64, element string) {
	log.Println("LRem ", h.key, " ", cnt, " ", element)
	cmd := h.cli.LRem(ctx, h.key, cnt, element)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}
