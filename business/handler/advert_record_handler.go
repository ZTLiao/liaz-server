package handler

import (
	"basic/device"
	"business/storage"
	"core/response"
	"core/web"
)

type AdvertRecordHandler struct {
	AdvertRecordDb *storage.AdvertRecordDb
}

func (e *AdvertRecordHandler) Record(wc *web.WebContext) interface{} {
	advertType := wc.PostForm("advertType")
	content := wc.PostForm("content")
	deviceInfo := device.GetDeviceInfo(wc)
	e.AdvertRecordDb.InsertAdvertRecord(advertType, content, deviceInfo)
	return response.Success()
}
