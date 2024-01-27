package handler

import (
	"business/storage"
	"core/response"
	"core/web"
)

type AppVersionHandler struct {
	AppVersionDb *storage.AppVersionDb
}

func (e *AppVersionHandler) CheckUpdate(wc *web.WebContext) interface{} {
	return response.Success()
}
