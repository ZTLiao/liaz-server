package controller

import (
	"admin/handler"
	"core/web"
)

type AdminUserController struct {
	handler.AdminUserHandler
}

func (e *AdminUserController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/user/get", e.GetUser)
}
