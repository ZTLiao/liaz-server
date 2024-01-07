package controller

import (
	"business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type SearchController struct {
}

var _ web.IWebController = &SearchController{}

func (e *SearchController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var searchHandler = &handler.SearchHandler{
		SearchDb:    storage.NewSearchDb(db),
		AssetDb:     storage.NewAssetDb(db),
		SearchCache: storage.NewSearchCache(redis),
	}
	iWebRoutes.GET("/search", searchHandler.Search)
	iWebRoutes.GET("/search/hot/rank", searchHandler.HotRank)
}
