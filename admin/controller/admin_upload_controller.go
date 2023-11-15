package controller

import (
	"admin/handler"
	"basic/infrastructure/persistence"
	"basic/interfaces"
	"core/application"
	"core/file"
	"core/web"
)

type AdminUploadController struct {
}

func (e *AdminUploadController) Router(iWebRoutes web.IWebRoutes) {
	var adminUploadHandler = &handler.AdminUploadHandler{
		FileItem: *interfaces.NewFileItem(
			persistence.NewFileItemRepository(application.GetXormEngine()),
			*file.NewFileTemplate()),
	}
	iWebRoutes.POST("/upload/:bucketName", adminUploadHandler.Upload)
}
