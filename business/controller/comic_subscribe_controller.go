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

type ComicSubscribeController struct {
}

var _ web.IWebController = &ComicSubscribeController{}

func (e *ComicSubscribeController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var comicSubscribeNumCache = storage.NewComicSubscribeNumCache(redis)
	event.Bus.Subscribe(constant.COMIC_SUBSCRIBE_TOPIC, listener.NewComicSubscribeListener(storage.NewComicDb(db), comicSubscribeNumCache))
	event.Bus.Subscribe(constant.COMIC_SUBSCRIBE_RANK_TOPIC, listener.NewComicSubscribeRankListener(storage.NewComicRankCache(redis)))
	var comicSubscribeHandler = &handler.ComicSubscribeHandler{
		ComicSubscribeDb:       storage.NewComicSubscribeDb(db),
		ComicSubscribeNumCache: comicSubscribeNumCache,
	}
	iWebRoutes.POST("/comic/subscribe", comicSubscribeHandler.Subscribe)
}
