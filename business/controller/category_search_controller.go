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
	var comicHandler = &handler.ComicHandler{
		ComicDb:            storage.NewComicDb(db),
		ComicChapterDb:     storage.NewComicChapterDb(db),
		ComicChapterItemDb: storage.NewComicChapterItemDb(db),
	}
	var categorySearchHandler = &handler.CategorySearchHandler{
		ComicHandler: comicHandler,
	}
	iWebRoutes.GET("/category/search", categorySearchHandler.GetContent)
}
