package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AppVersionController struct {
}

var _ web.IWebController = &AppVersionController{}

func (e *AppVersionController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var appVersionHandler = handler.AppVersionHandler{
		AppVersionDb: storage.NewAppVersionDb(db),
	}
	iWebRoutes.GET("/app/version/check/update", appVersionHandler.CheckUpdate)
}
