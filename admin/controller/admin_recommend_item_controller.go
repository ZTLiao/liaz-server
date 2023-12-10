package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminRecommendItemController struct {
}

var _ web.IWebController = &AdminRecommendItemController{}

func (e *AdminRecommendItemController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminRecommendItemHandler = &handler.AdminRecommendItemHandler{
		RecommendItemDb: storage.NewRecommendItemDb(db),
	}
	iWebRoutes.GET("/recommend/item/page", adminRecommendItemHandler.GetRecommendItemPage)
	iWebRoutes.POST("/recommend/item", adminRecommendItemHandler.SaveRecommendItem)
	iWebRoutes.PUT("/recommend/item", adminRecommendItemHandler.UpdateRecommendItem)
	iWebRoutes.DELETE("/recommend/item/:recommendItemId", adminRecommendItemHandler.DelRecommendItem)
}
