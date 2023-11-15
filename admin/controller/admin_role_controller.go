package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminRoleController struct {
}

func (e *AdminRoleController) Router(iWebRoutes web.IWebRoutes) {
	var adminRoleHandler = &handler.AdminRoleHandler{
		AdminRoleDb:     storage.AdminRoleDb{},
		AdminRoleMenuDb: storage.AdminRoleMenuDb{},
	}
	iWebRoutes.GET("/role", adminRoleHandler.GetAdminRole)
	iWebRoutes.POST("/role", adminRoleHandler.SaveAdminRole)
	iWebRoutes.PUT("/role", adminRoleHandler.UpdateAdminRole)
	iWebRoutes.DELETE("/role/:roleId", adminRoleHandler.DelAdminRole)
}
