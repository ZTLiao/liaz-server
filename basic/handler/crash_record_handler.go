package handler

import (
	"basic/device"
	"basic/storage"
	"core/response"
	"core/web"
)

type CrashRecordHandler struct {
	CrashRecordDb *storage.CrashRecordDb
}

func (e *CrashRecordHandler) Report(wc *web.WebContext) interface{} {
	err := wc.PostForm("err")
	stackTrace := wc.PostForm("stackTrace")
	deviceInfo := device.GetDeviceInfo(wc)
	e.CrashRecordDb.InsertCrashRecord(err, stackTrace, deviceInfo)
	return response.Success()
}
