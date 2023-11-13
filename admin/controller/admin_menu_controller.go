package controller

import (
	"admin/handler"
	"core/web"
)

type AdminMenuController struct {
}

func (e *AdminMenuController) Router(iWebRoutes web.IWebRoutes) {
	iWebRoutes.GET("/menu/list", new(handler.AdminMenuHandler).GetAdminMenuList)
	iWebRoutes.GET("/menu", new(handler.AdminMenuHandler).GetAdminMenu)
	iWebRoutes.POST("/menu", new(handler.AdminMenuHandler).SaveAdminMenu)
	iWebRoutes.PUT("/menu", new(handler.AdminMenuHandler).UpdateAdminMenu)
	iWebRoutes.DELETE("/menu/:menuId", new(handler.AdminMenuHandler).DelAdminMenu)
}
