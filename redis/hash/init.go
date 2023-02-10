package hash

import (
	"context"
	"github.com/shzgithub2018/learn/redis/dal"
)

var ctx context.Context

func init() {
	ctx = context.Background()
	dal.InitRedis(ctx)
}
