package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminComicController struct {
}

var _ web.IWebController = &AdminComicController{}

func (e *AdminComicController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminComicHandler = handler.AdminComicHandler{
		ComicDb: storage.NewComicDb(db),
	}
	iWebRoutes.GET("/comic/page", adminComicHandler.GetComicPage)
}
