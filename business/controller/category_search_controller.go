package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type CategorySearchController struct {
}

var _ web.IWebController = &CategorySearchController{}

func (e *CategorySearchController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var categorySearchHandler = &handler.CategorySearchHandler{
		AssetDb: storage.NewAssetDb(db),
	}
	iWebRoutes.GET("/category/search", categorySearchHandler.GetContent)
}
