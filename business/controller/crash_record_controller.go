package controller

import (
	"basic/handler"
	"basic/storage"
	"core/system"
	"core/web"
)

type CrashRecordController struct {
}

var _ web.IWebController = &CrashRecordController{}

func (e *CrashRecordController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var crashRecordHandler = handler.CrashRecordHandler{
		CrashRecordDb: storage.NewCrashRecordDb(db),
	}
	iWebRoutes.POST("/crash/record/report", crashRecordHandler.Report)
}
