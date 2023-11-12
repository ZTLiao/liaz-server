package controller

import (
	"admin/handler"
	"core/web"
)

type AdminRoleController struct {
}

func (e *AdminRoleController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/role", new(handler.AdminRoleHandler).GetAdminRole)
	iWebRoutes.POST("/role", new(handler.AdminRoleHandler).SaveAdminRole)
	iWebRoutes.PUT("/role", new(handler.AdminRoleHandler).UpdateAdminRole)
	iWebRoutes.DELETE("/role/:roleId", new(handler.AdminRoleHandler).DelAdminRole)
}
