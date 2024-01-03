package controller

import (
	"business/handler"
	"business/storage"
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
	var comicSubscribeHandler = &handler.ComicSubscribeHandler{
		ComicSubscribeDb:       storage.NewComicSubscribeDb(db),
		ComicSubscribeNumCache: storage.NewComicSubscribeNumCache(redis),
	}
	iWebRoutes.POST("/comic/subscribe", comicSubscribeHandler.Subscribe)
}
