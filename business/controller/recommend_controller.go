package controller

import (
	basicHandler "basic/handler"
	basicStorage "basic/storage"
	businessHandler "business/handler"
	businessStorage "business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type RecommendController struct {
}

var _ web.IWebController = &RecommendController{}

func (e *RecommendController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var sysConfHandler = basicHandler.NewSysConfHandler(basicStorage.NewSysConfDb(db), basicStorage.NewSysConfCache(redis))
	var recommendHandler = &businessHandler.RecommendHandler{
		RecommendDb:     businessStorage.NewRecommendDb(db),
		RecommendItemDb: businessStorage.NewRecommendItemDb(db),
		RecommendCache:  businessStorage.NewRecommendCache(redis),
		AssetDb:         businessStorage.NewAssetDb(db),
		SysConfHandler:  sysConfHandler,
	}
	iWebRoutes.GET("/recommend/:position", recommendHandler.Recommend)
}
