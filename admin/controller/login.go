package controller

import (
	"admin/handler"
	"core/web"
)

type LoginController struct {
}

func (e *LoginController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/login", new(handler.LoginHandler).Login)
}
