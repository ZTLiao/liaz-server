package handler

import (
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"
	"net/http"
)

type AdminLogoutHandler struct {
	AdminUserCache   *storage.AdminUserCache
	AccessTokenCache *storage.AccessTokenCache
}

func (e *AdminLogoutHandler) Logout(wc *web.WebContext) interface{} {
	accessToken := wc.GetHeader(constant.AUTHORIZATION)
	adminUser, err := e.AdminUserCache.Get(accessToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	e.AccessTokenCache.Del(adminUser.AdminId)
	e.AdminUserCache.Del(accessToken)
	return response.Success()
}
