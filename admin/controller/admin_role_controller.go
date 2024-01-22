package controller

import (
	"admin/handler"
	"admin/storage"
	"core/system"
	"core/web"
)

type AdminRoleController struct {
}

var _ web.IWebController = &AdminRoleController{}

func (e *AdminRoleController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminRoleHandler = handler.AdminRoleHandler{
		AdminRoleDb:     storage.NewAdminRoleDb(db),
		AdminRoleMenuDb: storage.NewAdminRoleMenuDb(db),
	}
	iWebRoutes.GET("/role", adminRoleHandler.GetAdminRole)
	iWebRoutes.POST("/role", adminRoleHandler.SaveAdminRole)
	iWebRoutes.PUT("/role", adminRoleHandler.UpdateAdminRole)
	iWebRoutes.DELETE("/role/:roleId", adminRoleHandler.DelAdminRole)
}
