package main

import (
	"context"
	"github.com/shzgithub2018/learn/Redis/dal"
)

func main() {
	ctx := context.Background()
	dal.InitRedis(ctx)
}
