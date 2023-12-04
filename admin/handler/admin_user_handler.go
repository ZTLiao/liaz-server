package handler

import (
	"admin/model"
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"
	"fmt"
	"net/http"
	"strconv"
)

type AdminUserHandler struct {
	AdminUserDb        *storage.AdminUserDb
	AdminLoginRecordDb *storage.AdminLoginRecordDb
	AdminUserCache     *storage.AdminUserCache
	AccessTokenCache   *storage.AccessTokenCache
}

func (e *AdminUserHandler) GetAdminUser(wc *web.WebContext) interface{} {
	accessToken := wc.GetHeader(constant.AUTHORIZATION)
	adminUser, err := e.AdminUserCache.Get(accessToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	lastTime, err := e.AdminLoginRecordDb.GetLastTime(adminUser.AdminId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(&resp.AdminUserResp{
		AdminId:  adminUser.AdminId,
		Name:     adminUser.Name,
		Username: adminUser.Username,
		Avatar:   adminUser.Avatar,
		LastTime: lastTime,
	})
}

func (e *AdminUserHandler) GetAdminUserList(wc *web.WebContext) interface{} {
	adminUsers, err := e.AdminUserDb.GetAdminUserList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminUsers)
}

func (e *AdminUserHandler) SaveAdminUser(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminUser(wc)
	return response.Success()
}

func (e *AdminUserHandler) UpdateAdminUser(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminUser(wc)
	return response.Success()
}

func (e *AdminUserHandler) saveOrUpdateAdminUser(wc *web.WebContext) {
	var params map[string]any
	if err := wc.ShouldBindJSON(&params); err != nil {
		wc.AbortWithError(err)
	}
	adminIdStr := fmt.Sprint(params["adminId"])
	name := fmt.Sprint(params["name"])
	username := fmt.Sprint(params["username"])
	password := fmt.Sprint(params["password"])
	phone := fmt.Sprint(params["phone"])
	avatar := fmt.Sprint(params["avatar"])
	email := fmt.Sprint(params["email"])
	introduction := fmt.Sprint(params["introduction"])
	statusStr := fmt.Sprint(params["status"])
	var adminUser = new(model.AdminUser)
	if len(adminIdStr) > 0 {
		adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		adminUser.AdminId = adminId
	}
	adminUser.Name = name
	adminUser.Username = username
	adminUser.Password = password
	adminUser.Phone = phone
	adminUser.Avatar = avatar
	adminUser.Email = email
	adminUser.Introduction = introduction
	status, err := strconv.ParseInt(statusStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminUser.Status = int8(status)
	e.AdminUserDb.SaveOrUpdateAdminUser(adminUser)
}

func (e *AdminUserHandler) DelAdminUser(wc *web.WebContext) interface{} {
	adminIdStr := wc.Param("adminId")
	if len(adminIdStr) > 0 {
		adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AdminUserDb.DelAdminUser(adminId)
		accessToken, err := e.AccessTokenCache.Get(adminId)
		if err != nil {
			wc.AbortWithError(err)
		}
		if len(accessToken) > 0 {
			e.AdminUserCache.Del(accessToken)
		}
	}
	return response.Success()
}

func (e *AdminUserHandler) ThawAdminUser(wc *web.WebContext) interface{} {
	adminIdStr := wc.PostForm("adminId")
	if len(adminIdStr) > 0 {
		adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AdminUserDb.ThawAdminUser(adminId)
		accessToken, err := e.AccessTokenCache.Get(adminId)
		if err != nil {
			wc.AbortWithError(err)
		}
		if len(accessToken) > 0 {
			e.AdminUserCache.Del(accessToken)
		}
	}
	return response.Success()
}
