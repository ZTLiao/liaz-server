package controller

import (
	"business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type NovelSubscribeController struct {
}

var _ web.IWebController = &NovelSubscribeController{}

func (e *NovelSubscribeController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var novelSubscribeHandler = &handler.NovelSubscribeHandler{
		NovelSubscribeDb:       storage.NewNovelSubscribeDb(db),
		NovelSubscribeNumCache: storage.NewNovelSubscribeNumCache(redis),
	}
	iWebRoutes.POST("/novel/subscribe", novelSubscribeHandler.Subscribe)
}
