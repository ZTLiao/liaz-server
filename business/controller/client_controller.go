package controller

import (
	basic "basic/handler"
	"basic/storage"
	"business/handler"
	"business/listener"
	"core/config"
	"core/constant"
	"core/event"
	"core/redis"
	"core/system"
	"core/web"
)

type ClientController struct {
}

var _ web.IWebController = &ClientController{}

func (e *ClientController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var sysConfHandler = basic.NewSysConfHandler(storage.NewSysConfDb(db), storage.NewSysConfCache(redis))
	//事件订阅
	event.Bus.Subscribe(constant.CLIENT_INIT_TOPIC, listener.NewClientInitListener(storage.NewDeviceDb(db), storage.NewClientInitRecordDb(db)))
	var clientHandler = &handler.ClientHandler{
		SysConfHandler: sysConfHandler,
		SecurityConfig: config.SystemConfig.Security,
	}
	iWebRoutes.GET("/client/init", clientHandler.ClientInit)
}
