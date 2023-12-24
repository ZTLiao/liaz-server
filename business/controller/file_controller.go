package controller

import (
	basicHandler "basic/handler"
	"basic/storage"
	businessHandler "business/handler"
	"core/file"
	"core/system"
	"core/web"
)

type FileController struct {
}

var _ web.IWebController = &FileController{}

func (e *FileController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	minio := system.GetMinioClient()
	var fileItemDb = storage.NewFileItemDb(db)
	var fileTemplate = file.NewFileTemplate(minio)
	var fileHandler = &businessHandler.FileHandler{
		FileItemHandler: basicHandler.NewFileItemHandler(fileItemDb, fileTemplate),
	}
	iWebRoutes.POST("/file/upload", fileHandler.Upload)
}
