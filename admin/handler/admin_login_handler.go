package handler

import (
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/logger"
	"core/response"
	"core/web"

	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
)

type AdminLoginHandler struct {
	AdminUserDb        *storage.AdminUserDb
	AccessTokenCache   *storage.AccessTokenCache
	AdminUserCache     *storage.AdminUserCache
	AdminLoginRecordDb *storage.AdminLoginRecordDb
}

func (e *AdminLoginHandler) Login(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	password := wc.PostForm("password")
	wc.Info("username : %s, password : %s", username, password)
	adminUser, err := e.AdminUserDb.GetLoginUser(username, password)
	if err != nil {
		wc.AbortWithError(err)
	}
	if adminUser == nil {
		return response.Fail(constant.LOGIN_ERROR)
	}
	adminId := adminUser.AdminId
	var accessToken string
	ok, err := e.AccessTokenCache.IsExist(adminId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if ok {
		accessToken, err = e.AccessTokenCache.Get(adminId)
		if err != nil {
			wc.AbortWithError(err)
		}
	} else {
		adminUser.Password = ""
		uuid, err := uuid.NewV4()
		if err != nil {
			logger.Error(err.Error())
		}
		accessToken = uuid.String()
		e.AccessTokenCache.Set(adminId, accessToken)
		e.AdminUserCache.Set(accessToken, adminUser)
	}
	//记录
	e.AdminLoginRecordDb.AddRecord(adminId, wc.ClientIP(), wc.GetHeader(constant.USER_AGENT))
	expireAt, err := e.AccessTokenCache.TTL(adminId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(&resp.AccessTokenResp{
		AccessToken: accessToken,
		ExpireAt:    expireAt,
		AdminId:     adminId,
	})
}
