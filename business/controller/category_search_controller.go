package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	businessStorage "business/storage"
	"core/system"
	"core/web"
)

type CategorySearchController struct {
}

var _ web.IWebController = &CategorySearchController{}

func (e *CategorySearchController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var categorySearchHandler = handler.CategorySearchHandler{
		CategoryGroupDb: basicStorage.NewCategoryGroupDb(db),
		CategoryDb:      basicStorage.NewCategoryDb(db),
		AssetDb:         businessStorage.NewAssetDb(db),
	}
	iWebRoutes.GET("/category/search", categorySearchHandler.GetCategorySearch)
}
