package controller

import (
	"business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type DiscussThumbController struct {
}

var _ web.IWebController = &DiscussThumbController{}

func (e *DiscussThumbController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var discussThumbHandler = handler.DiscussThumbHandler{
		DiscussThumbDb:       storage.NewDiscussThumbDb(db),
		DiscussThumbNumCache: storage.NewDiscussThumbNumCache(redis),
	}
	iWebRoutes.POST("/discuss/thumb", discussThumbHandler.Thumb)
}
