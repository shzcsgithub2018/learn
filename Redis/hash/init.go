package hash

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
)

var ctx context.Context

func init() {
	ctx = context.Background()
	dal.InitRedis(ctx)
}
