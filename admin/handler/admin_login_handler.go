package handler

import (
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"

	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
)

type AdminLoginHandler struct {
	AdminUserDb        storage.AdminUserDb
	AccessTokenCache   storage.AccessTokenCache
	AdminUserCache     storage.AdminUserCache
	AdminLoginRecordDb storage.AdminLoginRecordDb
}

// @Summary 登录
// @title Swagger API
// @Tags 授权管理
// @description 登录接口
// @BasePath /admin/login
// @Produce json
// @Param username formData string true "账号"
// @Param password formData string true "密码"
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/login [post]
func (e *AdminLoginHandler) Login(wc *web.WebContext) interface{} {
	var username = wc.Context.PostForm("username")
	var password = wc.Context.PostForm("password")
	wc.Info("username : %s, password : %s", username, password)
	adminUser := e.AdminUserDb.GetLoginUser(wc.Background(), username, password)
	if adminUser == nil {
		return response.Fail(constant.LOGIN_ERROR)
	}
	var adminId = adminUser.AdminId
	var accessToken string
	if e.AccessTokenCache.IsExist(wc.Background(), adminId) {
		accessToken = e.AccessTokenCache.Get(wc.Background(), adminId)
	} else {
		adminUser.Password = ""
		var uuid, _ = uuid.NewV4()
		accessToken = uuid.String()
		e.AccessTokenCache.Set(wc.Background(), adminId, accessToken)
		e.AdminUserCache.Set(wc.Background(), accessToken, adminUser)
	}
	//记录
	e.AdminLoginRecordDb.AddRecord(wc.Background(), adminId, wc.Context.ClientIP(), wc.Context.Request.Header.Get(constant.USER_AGENT))
	return response.ReturnOK(&resp.AccessTokenResp{
		AccessToken: accessToken,
		ExpireAt:    e.AccessTokenCache.TTL(wc.Background(), adminId),
		AdminId:     adminId,
	})
}
