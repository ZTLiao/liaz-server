package controller

import (
	"business/handler"
	"business/storage"
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
	var recommendHandler = &handler.RecommendHandler{
		RecommendDb:     storage.NewRecommendDb(db),
		RecommendItemDb: storage.NewRecommendItemDb(db),
		RecommendCache:  storage.NewRecommendCache(redis),
	}
	iWebRoutes.GET("/recommend/:position", recommendHandler.Recommend)
}
