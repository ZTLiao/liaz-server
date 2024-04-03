package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminNovelController struct {
}

var _ web.IWebController = &AdminNovelController{}

func (e *AdminNovelController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminNovelHandler = handler.AdminNovelHandler{
		NovelDb: storage.NewNovelDb(db),
	}
	iWebRoutes.GET("/novel/page", adminNovelHandler.GetNovelPage)
}
