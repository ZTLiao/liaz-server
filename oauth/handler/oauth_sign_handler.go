package handler

import (
	"basic/device"
	basicEnums "basic/enums"
	"basic/model"
	"basic/storage"
	"context"
	"core/constant"
	"core/event"
	"core/response"
	"core/types"
	"core/web"
	oauthEnums "oauth/enums"
	"oauth/resp"
	"strconv"

	"golang.org/x/oauth2"
)

type OauthSignHandler struct {
	Oauth2Config         *oauth2.Config
	AccountDb            *storage.AccountDb
	AccountLoginRecordDb *storage.AccountLoginRecordDb
	UserDeviceDb         *storage.UserDeviceDb
	UserDb               *storage.UserDb
	OAuth2TokenCache     *storage.OAuth2TokenCache
}

func (e *OauthSignHandler) SignIn(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	password := wc.PostForm("password")
	grantType := wc.PostForm("grantType")
	var token *oauth2.Token
	var account *model.Account
	var err error
	if grantType == oauthEnums.GRANT_TYPE_FOR_PASSWORD {
		token, err = e.Oauth2Config.PasswordCredentialsToken(context.Background(), username, password)
		if err != nil {
			if e, ok := err.(*oauth2.RetrieveError); ok {
				wc.Error("errorCode : %s, errorDescription : %s", e.ErrorCode, e.ErrorDescription)
				return response.Fail(e.ErrorDescription)
			}
			wc.AbortWithError(err)
		}
		account, err = e.AccountDb.SignIn(username, password)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	userId := account.UserId
	//设备信息
	deviceInfo := device.GetDeviceInfo(wc)
	//更新用户IP地址
	e.UserDb.UpdateLocation(userId, deviceInfo.ClientIp)
	//更新设备信息
	e.UserDeviceDb.SaveOrUpdateUserDevice(userId, deviceInfo.DeviceId)
	//更新token
	e.OAuth2TokenCache.Set(userId, token.AccessToken)
	//保存登录记录
	record, _ := e.AccountLoginRecordDb.InsertAccountLoginRecord(userId, deviceInfo)
	//发布事件
	event.Bus.Publish(constant.USER_LOGIN_TOPIC, record)
	return response.ReturnOK(&resp.CustomOAuth2TokenResp{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       types.Time(token.Expiry),
		UserId:       userId,
	})
}

func (e *OauthSignHandler) SignUp(wc *web.WebContext) interface{} {
	password := wc.PostForm("password")
	nickname := wc.PostForm("nickname")
	avatar := wc.PostForm("avatar")
	genderStr := wc.PostForm("gender")
	grantType := wc.PostForm("grantType")
	var username string
	var flag int8
	var gender int64
	if grantType == oauthEnums.GRANT_TYPE_FOR_PASSWORD {
		username = wc.PostForm("username")
		if len(username) == 0 {
			return response.Fail(constant.USERNAME_EMPTY)
		}
		if len(password) == 0 {
			return response.Fail(constant.PASSWORD_EMPTY)
		}
		if len(nickname) == 0 {
			return response.Fail(constant.NICKNAME_EMPTY)
		}
		var err error
		if len(genderStr) > 0 {
			gender, err = strconv.ParseInt(genderStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
		}
		flag |= basicEnums.ACCOUNT_FLAG_FOR_USERNAME
	}
	userId, err := e.AccountDb.SignUpForUsername(username, password, flag)
	if err != nil {
		return response.Fail(err.Error())
	}
	err = e.UserDb.SignUp(userId, nickname, avatar, int8(gender), basicEnums.USER_TYPE_FOR_NORMAL)
	if err != nil {
		wc.AbortWithError(err)
	}
	return e.SignIn(wc)
}
