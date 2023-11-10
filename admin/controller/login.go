package controller

import (
	"admin/handler"
	"core/web"
)

type LoginController struct {
	handler.LoginHandler
}

func (e *LoginController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/login", e.Login)
}
