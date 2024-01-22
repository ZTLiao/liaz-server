package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type BookshelfController struct {
}

var _ web.IWebController = &BookshelfController{}

func (e *BookshelfController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var bookshelfHandler = handler.BookshelfHandler{
		ComicChapterDb: storage.NewComicChapterDb(db),
		NovelChapterDb: storage.NewNovelChapterDb(db),
	}
	iWebRoutes.GET("/bookshelf/comic", bookshelfHandler.GetComic)
	iWebRoutes.GET("/bookshelf/novel", bookshelfHandler.GetNovel)
}
