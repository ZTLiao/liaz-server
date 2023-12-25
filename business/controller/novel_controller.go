package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type NovelController struct {
}

var _ web.IWebController = &NovelController{}

func (e *NovelController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var novelHandler = handler.NovelHandler{
		NovelDb:            storage.NewNovelDb(db),
		NovelChapterDb:     storage.NewNovelChapterDb(db),
		NovelChapterItemDb: storage.NewNovelChapterItemDb(db),
	}
	iWebRoutes.GET("/novel/:novelId", novelHandler.NovelDetail)
	iWebRoutes.GET("/novel/upgrade", novelHandler.NovelUpgrade)
}
