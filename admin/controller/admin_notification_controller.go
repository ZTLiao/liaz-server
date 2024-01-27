package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminNotificationController struct {
}

var _ web.IWebController = &AdminNotificationController{}

func (e *AdminNotificationController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminNotificationHandler = handler.AdminNotificationHandler{
		NotificationDb: storage.NewNotificationDb(db),
	}
	iWebRoutes.GET("/notification/page", adminNotificationHandler.GetNotificationPage)
	iWebRoutes.POST("/notification", adminNotificationHandler.SaveNotification)
	iWebRoutes.PUT("/notification", adminNotificationHandler.UpdateNotification)
	iWebRoutes.DELETE("/notification/:notificationId", adminNotificationHandler.DelNotification)
}
