package handler

import (
	"admin/resp"
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"

	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
)

type LoginHandler struct {
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
func (e *LoginHandler) Login(wc *web.WebContext) interface{} {
	var username = wc.Context.PostForm("username")
	var password = wc.Context.PostForm("password")
	wc.Info("username : %s, password : %s", username, password)
	adminUser := new(storage.AdminUserDb).GetLoginUser(username, password)
	if adminUser == nil {
		return response.Fail(constant.LOGIN_ERROR)
	}
	var adminId = adminUser.AdminId
	var accessTokenCache = new(storage.AccessTokenCache)
	var accessToken string
	if accessTokenCache.IsExist(adminId) {
		accessToken = accessTokenCache.Get(adminId)
	} else {
		adminUser.Password = ""
		var uuid, _ = uuid.NewV4()
		accessToken = uuid.String()
		accessTokenCache.Set(adminId, accessToken)
		new(storage.AdminUserCache).Set(accessToken, adminUser)
	}
	//记录
	new(AdminLoginRecordHandler).AddRecord(adminId, wc.Context.ClientIP(), wc.Context.Request.Header.Get(constant.USER_AGENT))
	return response.ReturnOK(&resp.AccessTokenResp{
		AccessToken: accessToken,
		ExpireAt:    accessTokenCache.TTL(adminId),
		AdminId:     adminId,
	})
}
