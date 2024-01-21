package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type MessageBoardController struct {
}

var _ web.IWebController = &MessageBoardController{}

func (e *MessageBoardController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var messageBoardHandler = handler.MessageBoardHandler{
		MessageBoardDb: storage.NewMessageBoardDb(db),
	}
	iWebRoutes.POST("/message/board/welcome", messageBoardHandler.Welcome)
}
