package controller

import (
	"admin/handler"
	"core/web"
)

type AdminLoginController struct {
}

func (e *AdminLoginController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/login", new(handler.AdminLoginHandler).Login)
}
