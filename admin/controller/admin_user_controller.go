package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminUserController struct {
}

func (e *AdminUserController) Router(iWebRoutes web.IWebRoutes) {
	var adminUserHandler = &handler.AdminUserHandler{
		AdminUserDb:        storage.AdminUserDb{},
		AdminLoginRecordDb: storage.AdminLoginRecordDb{},
		AdminUserCache:     storage.AdminUserCache{},
		AccessTokenCache:   storage.AccessTokenCache{},
	}
	iWebRoutes.GET("/user/get", adminUserHandler.GetAdminUser)
	iWebRoutes.GET("/user", adminUserHandler.GetAdminUserList)
	iWebRoutes.POST("/user", adminUserHandler.SaveAdminUser)
	iWebRoutes.PUT("/user", adminUserHandler.UpdateAdminUser)
	iWebRoutes.DELETE("/user/:adminId", adminUserHandler.DelAdminUser)
	iWebRoutes.PUT("/user/thaw", adminUserHandler.ThawAdminUser)
}
