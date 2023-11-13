package controller

import (
	"admin/handler"
	"core/web"
)

type LogoutController struct {
}

func (e *LogoutController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/logout", new(handler.LogoutHandler).Logout)
}
