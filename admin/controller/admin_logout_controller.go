package controller

import (
	"admin/handler"
	"core/web"
)

type AdminLogoutController struct {
}

func (e *AdminLogoutController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.POST("/logout", new(handler.AdminLogoutHandler).Logout)
}
