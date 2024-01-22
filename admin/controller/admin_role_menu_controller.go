package controller

import (
	"admin/handler"
	"admin/storage"
	"core/system"
	"core/web"
)

type AdminRoleMenuController struct {
}

var _ web.IWebController = &AdminRoleMenuController{}

func (e *AdminRoleMenuController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminRoleMenuHandler = handler.AdminRoleMenuHandler{
		AdminRoleMenuDb: storage.NewAdminRoleMenuDb(db),
		AdminMenuDb:     storage.NewAdminMenuDb(db),
	}
	iWebRoutes.GET("/role/menu/:roleId", adminRoleMenuHandler.GetAdminRoleMenu)
	iWebRoutes.POST("/role/menu", adminRoleMenuHandler.SaveAdminRoleMenu)
}
