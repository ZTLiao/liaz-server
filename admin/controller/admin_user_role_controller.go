package controller

import (
	"admin/handler"
	"admin/storage"
	"core/system"
	"core/web"
)

type AdminUserRoleController struct {
}

var _ web.IWebController = &AdminUserRoleController{}

func (e *AdminUserRoleController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminUserRoleHandler = handler.AdminUserRoleHandler{
		AdminUserRoleDb: storage.NewAdminUserRoleDb(db),
		AdminRoleDb:     storage.NewAdminRoleDb(db),
	}
	iWebRoutes.GET("/user/role/:adminId", adminUserRoleHandler.GetAdminUserRole)
	iWebRoutes.POST("/user/role", adminUserRoleHandler.SaveAdminUserRole)
}
