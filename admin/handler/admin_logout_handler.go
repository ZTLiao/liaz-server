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

// @Summary 登出
// @title Swagger API
// @Tags 授权管理
// @description 登出接口
// @Security ApiKeyAuth
// @BasePath /admin/logout
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/logout [post]
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
