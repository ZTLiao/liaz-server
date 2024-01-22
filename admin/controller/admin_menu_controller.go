package controller

import (
	"admin/handler"
	"admin/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminMenuController struct {
}

var _ web.IWebController = &AdminMenuController{}

func (e *AdminMenuController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var adminMenuHandler = handler.AdminMenuHandler{
		AdminMenuDb:     storage.NewAdminMenuDb(db),
		AdminRoleMenuDb: storage.NewAdminRoleMenuDb(db),
		AdminUserCache:  storage.NewAdminUserCache(redis),
	}
	iWebRoutes.GET("/menu/list", adminMenuHandler.GetAdminMenuList)
	iWebRoutes.GET("/menu", adminMenuHandler.GetAdminMenu)
	iWebRoutes.POST("/menu", adminMenuHandler.SaveAdminMenu)
	iWebRoutes.PUT("/menu", adminMenuHandler.UpdateAdminMenu)
	iWebRoutes.DELETE("/menu/:menuId", adminMenuHandler.DelAdminMenu)
}
