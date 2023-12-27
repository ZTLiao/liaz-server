package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	businessStorage "business/storage"
	"core/system"
	"core/web"
)

type NovelChapterController struct {
}

var _ web.IWebController = &NovelChapterController{}

func (e *NovelChapterController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var novelChapterHandler = &handler.NovelChapterHandler{
		NovelDb:            businessStorage.NewNovelDb(db),
		NovelChapterDb:     businessStorage.NewNovelChapterDb(db),
		NovelChapterItemDb: businessStorage.NewNovelChapterItemDb(db),
		FileItemDb:         basicStorage.NewFileItemDb(db),
	}
	iWebRoutes.GET("/novel/chapter", novelChapterHandler.GetNovelChapter)
}
