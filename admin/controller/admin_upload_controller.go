package controller

import (
	adminHandler "admin/handler"
	basicHandler "basic/handler"
	"basic/storage"
	"core/file"
	"core/system"
	"core/web"
)

type AdminUploadController struct {
}

var _ web.IWebController = &AdminUploadController{}

func (e *AdminUploadController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	minio := system.GetMinioClient()
	var fileItemDb = storage.NewFileItemDb(db)
	var fileTemplate = file.NewFileTemplate(minio)
	var adminUploadHandler = &adminHandler.AdminUploadHandler{
		FileItemHandler: basicHandler.NewFileItemHandler(fileItemDb, fileTemplate),
	}
	iWebRoutes.POST("/upload/:bucketName", adminUploadHandler.Upload)
}
