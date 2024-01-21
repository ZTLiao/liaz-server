package handler

import (
	"basic/device"
	"business/storage"
	"core/response"
	"core/web"
)

type MessageBoardHandler struct {
	MessageBoardDb *storage.MessageBoardDb
}

func (e *MessageBoardHandler) Welcome(wc *web.WebContext) interface{} {
	content := wc.PostForm("content")
	deviceInfo := device.GetDeviceInfo(wc)
	e.MessageBoardDb.InsertMessageBoard(content, deviceInfo)
	return response.Success()
}
