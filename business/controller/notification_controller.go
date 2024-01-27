package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type NotificationController struct {
}

var _ web.IWebController = &NotificationController{}

func (e *NotificationController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var notificationHandler = handler.NotificationHandler{
		NotificationDb: storage.NewNotificationDb(db),
	}
	iWebRoutes.GET("/notification/latest", notificationHandler.GetLatest)
}
