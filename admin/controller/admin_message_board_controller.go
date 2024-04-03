package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminMessageBoardController struct {
}

var _ web.IWebController = &AdminMessageBoardController{}

func (e *AdminMessageBoardController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminMessageBoardHandler = handler.AdminMessageBoardHandler{
		MessageBoardDb: storage.NewMessageBoardDb(db),
	}
	iWebRoutes.GET("/message/board/page", adminMessageBoardHandler.GetMessageBoardPage)
}
