package controller

import (
	"admin/handler"
	"admin/storage"
	"core/web"
)

type AdminLogoutController struct {
}

func (e *AdminLogoutController) Router(iWebRoutes web.IWebRoutes) {
	var adminLogoutHandler = &handler.AdminLogoutHandler{
		AdminUserCache:   storage.AdminUserCache{},
		AccessTokenCache: storage.AccessTokenCache{},
	}
	iWebRoutes.POST("/logout", adminLogoutHandler.Logout)
}
