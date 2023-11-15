package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminMenuController struct {
}

func (e *AdminMenuController) Router(iWebRoutes web.IWebRoutes) {
	var adminMenuHandler = &handler.AdminMenuHandler{
		AdminMenuDb:     storage.AdminMenuDb{},
		AdminRoleMenuDb: storage.AdminRoleMenuDb{},
		AdminUserCache:  storage.AdminUserCache{},
	}
	iWebRoutes.GET("/menu/list", adminMenuHandler.GetAdminMenuList)
	iWebRoutes.GET("/menu", adminMenuHandler.GetAdminMenu)
	iWebRoutes.POST("/menu", adminMenuHandler.SaveAdminMenu)
	iWebRoutes.PUT("/menu", adminMenuHandler.UpdateAdminMenu)
	iWebRoutes.DELETE("/menu/:menuId", adminMenuHandler.DelAdminMenu)
}
