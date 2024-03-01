package controller

import (
	"basic/storage"
	"business/handler"
	"core/system"
	"core/web"
)

type CategoryController struct {
}

var _ web.IWebController = &CategoryController{}

func (e *CategoryController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var categoryGroupHandler = handler.CategoryGroupHandler{
		CategoryGroupDb: storage.NewCategoryGroupDb(db),
	}
	iWebRoutes.GET("/category", categoryGroupHandler.GetGroupToCategory)
}
