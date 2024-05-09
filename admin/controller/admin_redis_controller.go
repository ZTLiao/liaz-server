package controller

import (
	"admin/handler"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminRedisController struct {
}

var _ web.IWebController = &AdminRedisController{}

func (e *AdminRedisController) Router(iWebRoutes web.IWebRoutes) {
	var adminRedisHandler = handler.AdminRedisHandler{
		RedisUtil: redis.NewRedisUtil(system.GetRedisClient()),
	}
	iWebRoutes.DELETE("/delete", adminRedisHandler.Delete)
}
