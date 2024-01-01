package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type ComicSubscribeController struct {
}

var _ web.IWebController = &ComicSubscribeController{}

func (e *ComicSubscribeController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var comicSubscribeHandler = &handler.ComicSubscribeHandler{
		ComicSubscribeDb: storage.NewComicSubscribeDb(db),
	}
	iWebRoutes.POST("/comic/subscribe", comicSubscribeHandler.Subscribe)
}
