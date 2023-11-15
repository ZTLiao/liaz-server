package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminUserRoleController struct {
}

func (e *AdminUserRoleController) Router(iWebRoutes web.IWebRoutes) {
	var adminUserRoleHandler = &handler.AdminUserRoleHandler{
		AdminUserRoleDb: storage.AdminUserRoleDb{},
		AdminRoleDb:     storage.AdminRoleDb{},
	}
	iWebRoutes.GET("/user/role/:adminId", adminUserRoleHandler.GetAdminUserRole)
	iWebRoutes.POST("/user/role", adminUserRoleHandler.SaveAdminUserRole)
}
