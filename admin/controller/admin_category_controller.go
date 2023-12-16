package controller

import (
	"admin/handler"
	"basic/storage"
	"core/system"
	"core/web"
)

type AdminCategoryController struct {
}

var _ web.IWebController = &AdminCategoryController{}

func (e *AdminCategoryController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminCategoryHandler = &handler.AdminCategoryHandler{
		CategoryDb: storage.NewCategoryDb(db),
	}
	iWebRoutes.GET("/category/page", adminCategoryHandler.GetCategoryPage)
	iWebRoutes.GET("/cateogry", adminCategoryHandler.GetCategoryList)
	iWebRoutes.POST("/category", adminCategoryHandler.SaveCategory)
	iWebRoutes.PUT("/category", adminCategoryHandler.UpdateCategory)
	iWebRoutes.DELETE("/category/:categoryId", adminCategoryHandler.DelCategory)
}
