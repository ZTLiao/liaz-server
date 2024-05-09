package router

import (
	"task/handler"

	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	// 添加定时任务
	c.AddFunc("*/5 * * * * ?", new(handler.HelloWorldHandler).Execute)
	// 启动定时任务
	c.Start()
}
