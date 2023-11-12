package controller

import (
	"admin/handler"
	"core/web"
)

type AdminUserController struct {
}

func (e *AdminUserController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/user/get", new(handler.AdminUserHandler).GetAdminUser)
	iWebRoutes.GET("/user", new(handler.AdminUserHandler).GetAdminUserList)
	iWebRoutes.POST("/user", new(handler.AdminUserHandler).SaveAdminUser)
	iWebRoutes.PUT("/user", new(handler.AdminUserHandler).UpdateAdminUser)
	iWebRoutes.DELETE("/user/:adminId", new(handler.AdminUserHandler).DelAdminUser)
}
