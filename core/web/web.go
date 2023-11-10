package web

import (
	"github.com/gin-gonic/gin"
)

var (
	routers = make([]func(*WebRouterGroup), 0)
)

// 添加路由
func AddRouter(handlers ...func(*WebRouterGroup)) {
	routers = append(routers, handlers...)
}

// 初始化路由
func InitRouter(engine *gin.Engine) {
	var wrg = new(WebRouterGroup)
	wrg.group = &engine.RouterGroup
	for _, h := range routers {
		h(wrg)
	}
}
