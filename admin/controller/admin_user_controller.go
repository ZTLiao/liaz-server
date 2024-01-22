package controller

import (
	"admin/handler"
	"admin/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminUserController struct {
}

var _ web.IWebController = &AdminUserController{}

func (e *AdminUserController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var adminUserHandler = handler.AdminUserHandler{
		AdminUserDb:        storage.NewAdminUserDb(db),
		AdminLoginRecordDb: storage.NewAdminLoginRecordDb(db),
		AdminUserCache:     storage.NewAdminUserCache(redis),
		AccessTokenCache:   storage.NewAccessTokenCache(redis),
	}
	iWebRoutes.GET("/user/get", adminUserHandler.GetAdminUser)
	iWebRoutes.GET("/user", adminUserHandler.GetAdminUserList)
	iWebRoutes.POST("/user", adminUserHandler.SaveAdminUser)
	iWebRoutes.PUT("/user", adminUserHandler.UpdateAdminUser)
	iWebRoutes.DELETE("/user/:adminId", adminUserHandler.DelAdminUser)
	iWebRoutes.PUT("/user/thaw", adminUserHandler.ThawAdminUser)
}
