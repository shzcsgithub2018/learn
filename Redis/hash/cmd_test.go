package hash

import (
	"github.com/shzgithub2018/learn/Redis/dal"
	"github.com/shzgithub2018/learn/Util/convert"
	"testing"
)

var (
	key     = "article::10086"
	valHash = map[interface{}]interface{}{
		"title":      "greeting",
		"content":    "hello world",
		"author":     "peter",
		"created_at": "1231313.12313",
	}
)

func TestHSet(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(valHash)...)
	hash.HSet(convert.MapToList(valHash)...)
	/*Output
	4  set success.
	0  set success.   4  update success.
	*/
}

func TestHSetNX(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(valHash)...)
	hash.HSetNX("title", "greeting")
	hash.HSetNX("image", "nil")

	hash1 := NewHash(ctx, t, key+"plus", dal.GetRedisCli(ctx))
	hash1.HSetNX("title", "greeting")
	/*Output
	4  set success.
	fail.
	success.

	success.
	*/
}

func TestHGet(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(valHash)...)

	for k, _ := range valHash {
		hash.HGet(k.(string))
	}

	hash.HGet("picture")
	hash1 := NewHash(ctx, t, key+"plus", dal.GetRedisCli(ctx))
	hash1.HGet("title")

	/*Output
	4  set success.
	greeting
	hello world
	peter
	1231313.12313
	redis: nil
	redis: nil
	*/
}
