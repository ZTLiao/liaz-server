package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type NovelSubscribeController struct {
}

var _ web.IWebController = &NovelSubscribeController{}

func (e *NovelSubscribeController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var novelSubscribeHandler = &handler.NovelSubscribeHandler{
		NovelSubscribeDb: storage.NewNovelSubscribeDb(db),
	}
	iWebRoutes.POST("/novel/subscribe", novelSubscribeHandler.Subscribe)
}
