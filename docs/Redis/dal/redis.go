package dal

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	redisCli *redis.Client
)

func GetRedisCli(ctx context.Context) *redis.Client {
	return redisCli.WithContext(ctx)
}

func InitRedis(ctx context.Context) {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
