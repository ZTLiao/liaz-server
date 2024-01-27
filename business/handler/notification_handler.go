package handler

import (
	"business/storage"
	"core/response"
	"core/web"
)

type NotificationHandler struct {
	NotificationDb *storage.NotificationDb
}

func (e *NotificationHandler) GetLatest(wc *web.WebContext) interface{} {
	notification, err := e.NotificationDb.GetLatest()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(notification)
}
