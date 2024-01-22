package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type ComicChapterController struct {
}

var _ web.IWebController = &ComicChapterController{}

func (e *ComicChapterController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var comicChapterHandler = handler.ComicChapterHandler{
		ComicDb:            storage.NewComicDb(db),
		ComicChapterDb:     storage.NewComicChapterDb(db),
		ComicChapterItemDb: storage.NewComicChapterItemDb(db),
	}
	iWebRoutes.GET("/comic/chapter", comicChapterHandler.GetComicChapter)
}
