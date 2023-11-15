package handler

import (
	"admin/model"
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/errors"
	"core/response"
	"core/web"
	"fmt"
	"net/http"
	"strconv"
)

type AdminUserHandler struct {
	AdminUserDb        storage.AdminUserDb
	AdminLoginRecordDb storage.AdminLoginRecordDb
	AdminUserCache     storage.AdminUserCache
	AccessTokenCache   storage.AccessTokenCache
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
	var adminUser = e.AdminUserCache.Get(accessToken)
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	return response.ReturnOK(&resp.AdminUserResp{
		AdminUser: *adminUser,
		LastTime:  e.AdminLoginRecordDb.GetLastTime(adminUser.AdminId),
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
	return response.ReturnOK(e.AdminUserDb.GetAdminUserList(wc.Background()))
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
	var params map[string]any
	if err := wc.Context.ShouldBindJSON(&params); err != nil {
		wc.Context.Error(&errors.ApiError{
			Message: err.Error(),
		})
		return
	}
	var adminIdStr = fmt.Sprint(params["adminId"])
	var name = fmt.Sprint(params["name"])
	var username = fmt.Sprint(params["username"])
	var password = fmt.Sprint(params["password"])
	var phone = fmt.Sprint(params["phone"])
	var avatar = fmt.Sprint(params["avatar"])
	var email = fmt.Sprint(params["email"])
	var introduction = fmt.Sprint(params["introduction"])
	var statusStr = fmt.Sprint(params["status"])
	var adminUser = new(model.AdminUser)
	if len(adminIdStr) > 0 {
		adminUser.AdminId, _ = strconv.ParseInt(adminIdStr, 10, 64)
	}
	adminUser.Name = name
	adminUser.Username = username
	adminUser.Password = password
	adminUser.Phone = phone
	adminUser.Avatar = avatar
	adminUser.Email = email
	adminUser.Introduction = introduction
	status, _ := strconv.ParseInt(statusStr, 10, 64)
	adminUser.Status = int8(status)
	e.AdminUserDb.SaveOrUpdateAdminUser(adminUser)
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
	var adminIdStr = wc.Context.Param("adminId")
	if len(adminIdStr) > 0 {
		adminId, _ := strconv.ParseInt(adminIdStr, 10, 64)
		e.AdminUserDb.DelAdminUser(adminId)
		accessToken := e.AccessTokenCache.Get(adminId)
		if len(accessToken) > 0 {
			e.AdminUserCache.Del(accessToken)
		}
	}
	return response.Success()
}

// @Summary 解冻系统用户
// @title Swagger API
// @Tags 用户管理
// @description 解冻系统用户接口
// @Security ApiKeyAuth
// @BasePath /admin/user/thaw
// @Param adminId formdata int64 true "用户ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user/:adminId [delete]
func (e *AdminUserHandler) ThawAdminUser(wc *web.WebContext) interface{} {
	var adminIdStr = wc.Context.PostForm("adminId")
	if len(adminIdStr) > 0 {
		adminId, _ := strconv.ParseInt(adminIdStr, 10, 64)
		e.AdminUserDb.ThawAdminUser(adminId)
		accessToken := e.AccessTokenCache.Get(adminId)
		if len(accessToken) > 0 {
			e.AdminUserCache.Del(accessToken)
		}
	}
	return response.Success()
}
