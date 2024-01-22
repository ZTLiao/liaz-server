package controller

import (
	"admin/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdminAppVersionController struct {
}

var _ web.IWebController = &AdminAppVersionController{}

func (e *AdminAppVersionController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminAppVersionHandler = handler.AdminAppVersionHandler{
		AppVersionDb: storage.NewAppVersionDb(db),
	}
	iWebRoutes.GET("/app/version/page", adminAppVersionHandler.GetAppVersionPage)
}
