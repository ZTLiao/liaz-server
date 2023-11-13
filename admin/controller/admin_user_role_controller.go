package controller

import (
	"admin/handler"
	"core/web"
)

type AdminUserRoleController struct {
}

func (e *AdminUserRoleController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/user/role/:adminId", new(handler.AdminUserRoleHandler).GetAdminUserRole)
	iWebRoutes.POST("/user/role", new(handler.AdminUserRoleHandler).SaveAdminUserRole)
}
