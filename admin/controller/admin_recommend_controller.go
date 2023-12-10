package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminRecommendController struct {
}

var _ web.IWebController = &AdminRecommendController{}

func (e *AdminRecommendController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminRecommendHandler = &handler.AdminRecommendHandler{
		RecommendDb: storage.NewRecommendDb(db),
	}
	iWebRoutes.GET("/recommend/page", adminRecommendHandler.GetRecommendPage)
	iWebRoutes.GET("/recommend", adminRecommendHandler.GetRecommendList)
	iWebRoutes.POST("/recommend", adminRecommendHandler.SaveRecommend)
	iWebRoutes.PUT("/recommend", adminRecommendHandler.UpdateRecommend)
	iWebRoutes.DELETE("/recommend/:recommendId", adminRecommendHandler.DelRecommend)

}
