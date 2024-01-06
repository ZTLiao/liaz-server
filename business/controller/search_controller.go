package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type SearchController struct {
}

var _ web.IWebController = &SearchController{}

func (e *SearchController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var searchHandler = &handler.SearchHandler{
		AssetDb: storage.NewAssetDb(db),
	}
	iWebRoutes.GET("/search", searchHandler.Search)
}
