package controller

import (
	basic "basic/handler"
	"basic/storage"
	"business/handler"
	"core/redis"
	"core/system"
	"core/web"
)

type ClientController struct {
}

func (e *ClientController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var sysConfHandler = basic.NewSysConfHandler(storage.NewSysConfDb(db), storage.NewSysConfCache(redis))
	var clientHandler = &handler.ClientHandler{
		SysConfHandler: sysConfHandler,
	}
	iWebRoutes.GET("/client/init", clientHandler.ClientInit)
}
