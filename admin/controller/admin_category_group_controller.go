package controller

import (
	"admin/handler"
	"basic/storage"
	"core/system"
	"core/web"
)

type AdminCategoryGroupController struct {
}

var _ web.IWebController = &AdminCategoryGroupController{}

func (e *AdminCategoryGroupController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminCategoryGroupHandler = handler.AdminCategoryGroupHandler{
		CategoryGroupDb: storage.NewCategoryGroupDb(db),
	}
	iWebRoutes.GET("/category/group/page", adminCategoryGroupHandler.GetCategoryGroupPage)
	iWebRoutes.GET("/cateogry/group", adminCategoryGroupHandler.GetCategoryGroupList)
	iWebRoutes.POST("/category/group", adminCategoryGroupHandler.SaveCategoryGroup)
	iWebRoutes.PUT("/category/group", adminCategoryGroupHandler.UpdateCategoryGroup)
	iWebRoutes.DELETE("/category/group/:categoryGroupId", adminCategoryGroupHandler.DelCategoryGroup)
}
