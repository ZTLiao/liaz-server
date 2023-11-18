package controller

import (
	"admin/handler"
	"basic/infrastructure/persistence/repository"
	"basic/infrastructure/persistence/repository/store"
	"basic/interfaces"
	"core/file"
	"core/system"
	"core/web"
)

type AdminUploadController struct {
}

func (e *AdminUploadController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	minio := system.GetMinioClient()
	var fileItemStore = store.NewFileItemStore(db)
	var fileItemRepo = repository.NewFileItemRepo(fileItemStore)
	var fileTemplate = file.NewFileTemplate(minio)
	var adminUploadHandler = &handler.AdminUploadHandler{
		FileItemHandler: interfaces.NewFileItemHandler(fileItemRepo, fileTemplate),
	}
	iWebRoutes.POST("/upload/:bucketName", adminUploadHandler.Upload)
}
