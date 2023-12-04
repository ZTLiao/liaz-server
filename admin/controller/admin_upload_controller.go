package controller

import (
	admin "admin/handler"
	basic "basic/handler"
	"basic/storage"
	"core/file"
	"core/system"
	"core/web"
)

type AdminUploadController struct {
}

func (e *AdminUploadController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	minio := system.GetMinioClient()
	var fileItemDb = storage.NewFileItemDb(db)
	var fileTemplate = file.NewFileTemplate(minio)
	var adminUploadHandler = &admin.AdminUploadHandler{
		FileItemHandler: basic.NewFileItemHandler(fileItemDb, fileTemplate),
	}
	iWebRoutes.POST("/upload/:bucketName", adminUploadHandler.Upload)
}
