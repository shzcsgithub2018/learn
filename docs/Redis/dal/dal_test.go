package dal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var ctx = context.Background()

func TestPing(t *testing.T) {
	InitRedis(ctx)
	rdb := GetRedisCli(ctx)

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	} else {
		t.Log("connecting")
	}
}

func TestVersion(t *testing.T) {
	InitRedis(ctx)
	rdb := GetRedisCli(ctx)

	if info, err := rdb.Info(ctx).Result(); err != nil {
		panic(err)
	} else {
		t.Log(info)
	}afasfa
}

func TestExampleClient(t *testing.T) {
	InitRedis(ctx)
	rdb := GetRedisCli(ctx)

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
