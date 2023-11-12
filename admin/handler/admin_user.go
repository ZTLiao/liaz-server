package handler

import (
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"
	"net/http"
	"strconv"
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
func (e *AdminUserHandler) GetAdminUser(wc *web.WebContext) interface{} {
	var accessToken = wc.Context.Request.Header.Get(constant.AUTHORIZATION)
	var adminUser = new(storage.AdminUserCache).Get(accessToken)
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	return response.ReturnOK(&resp.AdminUserResp{
		AdminUser: *adminUser,
		LastTime:  new(storage.AdminLoginRecordDb).GetLastTime(adminUser.AdminId),
	})
}

// @Summary 获取系统用户列表
// @title Swagger API
// @Tags 用户管理
// @description 获取系统用户列表接口
// @Security ApiKeyAuth
// @BasePath /admin/user
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user [get]
func (e *AdminUserHandler) GetAdminUserList(wc *web.WebContext) interface{} {
	return response.ReturnOK(new(storage.AdminUserDb).GetAdminUserList())
}

// @Summary 添加系统用户
// @title Swagger API
// @Tags 用户管理
// @description 添加系统用户接口
// @Security ApiKeyAuth
// @BasePath /admin/user
// @Param adminUser body model.AdminUser true "系统用户"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user [post]
func (e *AdminUserHandler) SaveAdminUser(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminUser(wc)
	return response.Success()
}

// @Summary 修改系统用户
// @title Swagger API
// @Tags 用户管理
// @description 修改系统用户接口
// @Security ApiKeyAuth
// @BasePath /admin/user
// @Param adminUser body model.AdminUser true "系统用户"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user [put]
func (e *AdminUserHandler) UpdateAdminUser(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminUser(wc)
	return response.Success()
}

func (e *AdminUserHandler) saveOrUpdateAdminUser(wc *web.WebContext) {

}

// @Summary 修改系统用户
// @title Swagger API
// @Tags 用户管理
// @description 修改系统用户接口
// @Security ApiKeyAuth
// @BasePath /admin/user/:adminId
// @Param adminId query int64 true "用户ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user/:adminId [delete]
func (e *AdminUserHandler) DelAdminUser(wc *web.WebContext) interface{} {
	var adminId = wc.Context.Param("adminId")
	if len(adminId) > 0 {
		val, _ := strconv.ParseInt(adminId, 10, 64)
		new(storage.AdminUserDb).DelAdminUser(val)
		accessToken := new(storage.AccessTokenCache).Get(val)
		if len(accessToken) > 0 {
			new(storage.AdminUserCache).Del(accessToken)
		}
	}
	return response.Success()
}
