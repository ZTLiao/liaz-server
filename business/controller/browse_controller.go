package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type BrowseController struct {
}

var _ web.IWebController = &BrowseController{}

func (e *BrowseController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var browseHandler = &handler.BrowseHandler{
		BrowseDb:         storage.NewBrowseDb(db),
		HistoryDb:        storage.NewHistoryDb(db),
		ComicSubscribeDb: storage.NewComicSubscribeDb(db),
		NovelSubscribeDb: storage.NewNovelSubscribeDb(db),
	}
	iWebRoutes.POST("/browse/history", browseHandler.BrowseHistory)
}
