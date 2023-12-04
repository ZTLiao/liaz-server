package handler

import (
	"basic/handler"
	"core/web"
)

type AdminUploadHandler struct {
	FileItemHandler *handler.FileItemHandler
}

func (e *AdminUploadHandler) Upload(wc *web.WebContext) interface{} {
	return e.FileItemHandler.UploadFile(wc)
}
