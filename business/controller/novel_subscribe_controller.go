package controller

import (
	"business/handler"
	"business/listener"
	"business/storage"
	"core/constant"
	"core/event"
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
	var novelSubscribeNumCache = storage.NewNovelSubscribeNumCache(redis)
	event.Bus.Subscribe(constant.NOVEL_SUBSCRIBE_TOPIC, listener.NewNovelSubscribeListener(storage.NewNovelDb(db), novelSubscribeNumCache))
	var novelSubscribeHandler = &handler.NovelSubscribeHandler{
		NovelSubscribeDb:       storage.NewNovelSubscribeDb(db),
		NovelSubscribeNumCache: novelSubscribeNumCache,
	}
	iWebRoutes.POST("/novel/subscribe", novelSubscribeHandler.Subscribe)
}
