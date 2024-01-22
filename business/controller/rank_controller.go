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
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var rankHandler = handler.RankHandler{
		ComicDb:            storage.NewComicDb(db),
		NovelDb:            storage.NewNovelDb(db),
		ComicRankCache:     storage.NewComicRankCache(redis),
		NovelRankCache:     storage.NewNovelRankCache(redis),
		ComicRankItemCache: storage.NewComicRankItemCache(redis),
		NovelRankItemCache: storage.NewNovelRankItemCache(redis),
	}
	iWebRoutes.GET("/rank", rankHandler.Rank)
}
