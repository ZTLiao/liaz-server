package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type DiscussThumbController struct {
}

var _ web.IWebController = &DiscussThumbController{}

func (e *DiscussThumbController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var discussThumbHandler = handler.DiscussThumbHandler{
		DiscussThumbDb: storage.NewDiscussThumbDb(db),
	}
	iWebRoutes.POST("/discuss/thumb", discussThumbHandler.Thumb)
}
