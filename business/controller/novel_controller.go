package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	businessStorage "business/storage"
	"core/system"
	"core/web"
)

type NovelController struct {
}

var _ web.IWebController = &NovelController{}

func (e *NovelController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var novelHandler = handler.NovelHandler{
		NovelDb:            businessStorage.NewNovelDb(db),
		NovelChapterDb:     businessStorage.NewNovelChapterDb(db),
		NovelChapterItemDb: businessStorage.NewNovelChapterItemDb(db),
		FileItemDb:         basicStorage.NewFileItemDb(db),
		NovelSubscribeDb:   businessStorage.NewNovelSubscribeDb(db),
	}
	iWebRoutes.GET("/novel/:novelId", novelHandler.NovelDetail)
	iWebRoutes.GET("/novel/upgrade", novelHandler.NovelUpgrade)
	iWebRoutes.GET("/novel/catalogue", novelHandler.NovelCatalogue)
	iWebRoutes.GET("/novel/get", novelHandler.GetNovel)
}
