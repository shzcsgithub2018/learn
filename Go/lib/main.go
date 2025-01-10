package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	expr, err := cronexpr.Parse("*/5 * * * *")
	if err != nil {
		fmt.Println("解析Cron表达式出错:", err)
		return
	}
	// 获取下一次任务执行时间，并打印语义化后的描述
	next, err := expr.Next(time.Now())
	if err != nil {
		fmt.Println("获取下一次执行时间出错:", err)
		return
	}
	fmt.Println("Cron表达式语义：每隔5分钟执行一次，下一次执行时间:", next)
}
