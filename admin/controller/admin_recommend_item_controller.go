package controller

import (
	adminHandler "admin/handler"
	businessHandler "business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminRecommendItemController struct {
}

var _ web.IWebController = &AdminRecommendItemController{}

func (e *AdminRecommendItemController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var recommendHandler = &businessHandler.RecommendHandler{
		RecommendDb:     storage.NewRecommendDb(db),
		RecommendItemDb: storage.NewRecommendItemDb(db),
		RecommendCache:  storage.NewRecommendCache(redis),
	}
	var adminRecommendItemHandler = &adminHandler.AdminRecommendItemHandler{
		RecommendItemDb:  storage.NewRecommendItemDb(db),
		RecommendHandler: recommendHandler,
	}
	iWebRoutes.GET("/recommend/item/page", adminRecommendItemHandler.GetRecommendItemPage)
	iWebRoutes.POST("/recommend/item", adminRecommendItemHandler.SaveRecommendItem)
	iWebRoutes.PUT("/recommend/item", adminRecommendItemHandler.UpdateRecommendItem)
	iWebRoutes.DELETE("/recommend/item/:recommendItemId", adminRecommendItemHandler.DelRecommendItem)
}
