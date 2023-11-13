package controller

import (
	"admin/handler"
	"core/web"
)

type AdminRoleMenuController struct {
}

func (e *AdminRoleMenuController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/role/menu/:roleId", new(handler.AdminRoleMenuHandler).GetAdminRoleMenu)
	iWebRoutes.POST("/role/menu", new(handler.AdminRoleMenuHandler).SaveAdminRoleMenu)
}
