package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	"business/listener"
	businessStorage "business/storage"
	"core/constant"
	"core/event"
	"core/redis"
	"core/system"
	"core/web"
)

type DiscussController struct {
}

var _ web.IWebController = &DiscussController{}

func (e *DiscussController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	event.Bus.Subscribe(constant.COMIC_DISCUSS_RANK_TOPIC, listener.NewComicDiscussRankListener(businessStorage.NewComicRankCache(redis)))
	event.Bus.Subscribe(constant.NOVEL_DISCUSS_RANK_TOPIC, listener.NewNovelDiscussRankListener(businessStorage.NewNovelRankCache(redis)))
	var discussHandler = handler.DiscussHandler{
		DiscussDb:            businessStorage.NewDiscussDb(db),
		DiscussResourceDb:    businessStorage.NewDiscussResourceDb(db),
		UserDb:               basicStorage.NewUserDb(db),
		DiscussNumCache:      businessStorage.NewDiscussNumCache(redis),
		DiscussThumbNumCache: businessStorage.NewDiscussThumbNumCache(redis),
	}
	iWebRoutes.POST("/discuss", discussHandler.Discuss)
	iWebRoutes.GET("/discuss/page", discussHandler.GetDiscussPage)
}
