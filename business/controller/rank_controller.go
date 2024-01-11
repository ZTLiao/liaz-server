package controller

import (
	"business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type RankController struct {
}

var _ web.IWebController = &RankController{}

func (e *RankController) Router(iWebRoutes web.IWebRoutes) {
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var rankHandler = &handler.RankHandler{
		ComicRankCache: storage.NewComicRankCache(redis),
		NovelRankCache: storage.NewNovelRankCache(redis),
	}
	iWebRoutes.GET("/rank", rankHandler.Rank)
}
