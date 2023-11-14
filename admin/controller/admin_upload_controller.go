package controller

import (
	"admin/handler"
	"core/web"
)

type AdminUploadController struct {
}

func (e *AdminUploadController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/upload/:bucketName", new(handler.AdminUploadHandler).Upload)
}
