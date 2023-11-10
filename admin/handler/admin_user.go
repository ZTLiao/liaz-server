package handler

import (
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"
	"net/http"
)

type AdminUserHandler struct {
}

// @Summary 获取当前用户信息
// @title Swagger API
// @Tags 用户管理
// @description 获取当前用户信息接口
// @Security ApiKeyAuth
// @BasePath /admin/user/get
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user/get [get]
func (e *AdminUserHandler) GetUser(wc *web.WebContext) interface{} {
	var accessToken = wc.Context.Request.Header.Get(constant.AUTHORIZATION)
	var adminUser = new(storage.AdminUserCache).Get(accessToken)
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	return response.ReturnOK(adminUser)
}
