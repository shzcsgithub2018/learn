package hash

import (
	"github.com/shzgithub2018/learn/Redis/dal"
	"github.com/shzgithub2018/learn/Util/convert"
	"testing"
)

var (
	key     = "article::10086"
	valKeys = []string{"title", "content", "author", "created_at"}
	valHash = map[any]any{
		"title":      "greeting",
		"content":    "hello world",
		"author":     "peter",
		"created_at": "1231313.12313",
	}
	numValHash = map[any]any{
		"view_count": "30",
		"created_at": "313.12313",
		"title":      "greeting",
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

func TestHIncrBy(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(numValHash)...)

	n := int64(10)
	for k, _ := range numValHash {
		hash.HIncrBy(k.(string), n)
	}
	/*Output
	3  set success.
	40
	ERR hash value is not an integer
	ERR hash value is not an integer
	*/
}

func TestHIncrByFloat(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(numValHash)...)

	n := 3.1415926535
	for k, _ := range numValHash {
		hash.HIncrByFloat(k.(string), n)
	}
	/*Output
	3  set success.
	33.1415926535
	316.2647226535
	ERR hash value is not a float
	*/
}

func TestHExists(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))

	hash.HExists("view_count")
	hash.HSet(convert.MapToList(numValHash)...)
	hash.HExists("view_count")
	/*Output
	HExists article::10086 view_count
	NO
	3  set success.
	HExists article::10086 view_count
	Yes
	*/
}

func TestHDel(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))

	hash.HDel("view_count")
	hash.HSet(convert.MapToList(numValHash)...)
	hash.HExists("view_count")
	hash.HDel("view_count")
	hash.HExists("view_count")
	/*Output
	HDel article::10086
	view_count
	0 del success.
	3  set success.
	HExists article::10086 view_count
	Yes
	HDel article::10086
	view_count
	1 del success.
	HExists article::10086 view_count
	NO
	*/
}

func TestHLen(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))

	hash.HLen()
	hash.HSet(convert.MapToList(numValHash)...)
	hash.HLen()
	/*Output
	HLen article::10086
	0
	3  set success.
	HLen article::10086
	3
	*/
}

func TestHMGetHkeysHValsHGetAll(t *testing.T) {
	hash := NewHash(ctx, t, key, dal.GetRedisCli(ctx))
	hash.HSet(convert.MapToList(valHash)...)

	hash.HMGet(valKeys...)
	hash.HKeys()
	hash.HVals()
	hash.HGetAll()
	/*Output
	4  set success.
	title : greeting
	content : hello world
	author : peter
	created_at : 1231313.12313
	[title content author created_at]
	[greeting hello world peter 1231313.12313]
	map[author:peter content:hello world created_at:1231313.12313 title:greeting]
	*/
}
