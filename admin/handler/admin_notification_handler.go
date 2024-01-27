package handler

import (
	"admin/resp"
	"business/model"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminNotificationHandler struct {
	NotificationDb *storage.NotificationDb
}

func (e *AdminNotificationHandler) GetNotificationPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.NotificationDb.GetNotificationPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminNotificationHandler) SaveNotification(wc *web.WebContext) interface{} {
	e.saveOrUpdateNotification(wc)
	return response.Success()
}

func (e *AdminNotificationHandler) UpdateNotification(wc *web.WebContext) interface{} {
	e.saveOrUpdateNotification(wc)
	return response.Success()
}

func (e *AdminNotificationHandler) saveOrUpdateNotification(wc *web.WebContext) {
	var Notification = new(model.Notification)
	if err := wc.ShouldBindJSON(&Notification); err != nil {
		wc.AbortWithError(err)
	}
	e.NotificationDb.SaveOrUpdateNotification(Notification)
}

func (e *AdminNotificationHandler) DelNotification(wc *web.WebContext) interface{} {
	notificationIdStr := wc.Param("notificationId")
	if len(notificationIdStr) > 0 {
		notificationId, err := strconv.ParseInt(notificationIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.NotificationDb.DelNotification(notificationId)
	}
	return response.Success()
}
