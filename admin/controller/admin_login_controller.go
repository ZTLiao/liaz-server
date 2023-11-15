package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminLoginController struct {
}

func (e *AdminLoginController) Router(iWebRoutes web.IWebRoutes) {
	var adminLoginHandler = &handler.AdminLoginHandler{
		AdminUserDb:        storage.AdminUserDb{},
		AccessTokenCache:   storage.AccessTokenCache{},
		AdminUserCache:     storage.AdminUserCache{},
		AdminLoginRecordDb: storage.AdminLoginRecordDb{},
	}
	iWebRoutes.POST("/login", adminLoginHandler.Login)
}
