package controller

import (
	basicHandler "basic/handler"
	basicStorage "basic/storage"
	"business/handler"
	businessStorage "business/storage"
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
	var sysConfHandler = basicHandler.NewSysConfHandler(basicStorage.NewSysConfDb(db), basicStorage.NewSysConfCache(redis))
	var searchHandler = handler.SearchHandler{
		SysConfHandler: sysConfHandler,
		SearchDb:       businessStorage.NewSearchDb(db),
		AssetDb:        businessStorage.NewAssetDb(db),
		SearchCache:    businessStorage.NewSearchCache(redis),
	}
	iWebRoutes.GET("/search", searchHandler.Search)
	iWebRoutes.GET("/search/hot/rank", searchHandler.HotRank)
}
