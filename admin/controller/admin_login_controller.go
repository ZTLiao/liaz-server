package controller

import (
	"admin/handler"
	"admin/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminLoginController struct {
}

var _ web.IWebController = &AdminLoginController{}

func (e *AdminLoginController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var adminLoginHandler = handler.AdminLoginHandler{
		AdminUserDb:        storage.NewAdminUserDb(db),
		AccessTokenCache:   storage.NewAccessTokenCache(redis),
		AdminUserCache:     storage.NewAdminUserCache(redis),
		AdminLoginRecordDb: storage.NewAdminLoginRecordDb(db),
	}
	iWebRoutes.POST("/login", adminLoginHandler.Login)
}
