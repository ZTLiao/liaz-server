package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminRoleMenuController struct {
}

func (e *AdminRoleMenuController) Router(iWebRoutes web.IWebRoutes) {
	var adminRoleMenuHandler = &handler.AdminRoleMenuHandler{
		AdminRoleMenuDb: storage.AdminRoleMenuDb{},
		AdminMenuDb:     storage.AdminMenuDb{},
	}
	iWebRoutes.GET("/role/menu/:roleId", adminRoleMenuHandler.GetAdminRoleMenu)
	iWebRoutes.POST("/role/menu", adminRoleMenuHandler.SaveAdminRoleMenu)
}
