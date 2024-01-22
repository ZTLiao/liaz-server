package controller

import (
	"admin/handler"
	"admin/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminLogoutController struct {
}

var _ web.IWebController = &AdminLogoutController{}

func (e *AdminLogoutController) Router(iWebRoutes web.IWebRoutes) {
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var adminLogoutHandler = handler.AdminLogoutHandler{
		AdminUserCache:   storage.NewAdminUserCache(redis),
		AccessTokenCache: storage.NewAccessTokenCache(redis),
	}
	iWebRoutes.POST("/logout", adminLogoutHandler.Logout)
}
