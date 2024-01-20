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
	"core/utils"
	"core/web"
	"net/http"
	oauthEnums "oauth/enums"
	"oauth/resp"
	"strconv"
	"strings"

	"github.com/go-oauth2/redis/v4"
	"golang.org/x/oauth2"
)

type OAuthSignHandler struct {
	OAuth2Config         *oauth2.Config
	AccountDb            *storage.AccountDb
	AccountLoginRecordDb *storage.AccountLoginRecordDb
	UserDeviceDb         *storage.UserDeviceDb
	UserDb               *storage.UserDb
	OAuth2TokenCache     *storage.OAuth2TokenCache
	RedisTokenStore      *redis.TokenStore
}

func (e *OAuthSignHandler) SignIn(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	password := wc.PostForm("password")
	grantType := wc.PostForm("grantType")
	var token *oauth2.Token
	var account *model.Account
	var err error
	if grantType == oauthEnums.GRANT_TYPE_FOR_PASSWORD {
		token, err = e.OAuth2Config.PasswordCredentialsToken(context.Background(), username, password)
		if err != nil {
			if e, ok := err.(*oauth2.RetrieveError); ok {
				wc.Error("errorCode : %s, errorDescription : %s", e.ErrorCode, e.ErrorDescription)
				return response.Fail(constant.USERNAME_OR_PASSWORD_ERROR)
			}
			wc.AbortWithError(err)
		}
		account, err = e.AccountDb.SignIn(username, password)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	if account == nil {
		return response.ReturnError(http.StatusForbidden, constant.USERNAME_NOT_EXISTS)
	}
	userId := account.UserId
	//设备信息
	deviceInfo := device.GetDeviceInfo(wc)
	//更新用户IP地址
	err = e.UserDb.UpdateLocation(userId, deviceInfo.ClientIp)
	if err != nil {
		wc.AbortWithError(err)
	}
	//更新设备信息
	err = e.UserDeviceDb.SaveOrUpdateUserDevice(userId, deviceInfo.DeviceId)
	if err != nil {
		wc.AbortWithError(err)
	}
	//更新token
	err = e.OAuth2TokenCache.Set(userId, token.AccessToken)
	if err != nil {
		wc.AbortWithError(err)
	}
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

func (e *OAuthSignHandler) SignUp(wc *web.WebContext) interface{} {
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

func (e *OAuthSignHandler) SignOut(wc *web.WebContext) interface{} {
	clientToken := wc.GetHeader(constant.AUTHORIZATION)
	if len(clientToken) == 0 {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	tokenArray := strings.Split(clientToken, utils.SPACE)
	tokenType := tokenArray[0]
	if tokenType != constant.TOKEN_TYPE {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	clientToken = tokenArray[1]
	userIdStr := wc.GetHeader(constant.X_USER_ID)
	if len(userIdStr) == 0 {
		return response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED)
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	serverToken, err := e.OAuth2TokenCache.Get(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if clientToken != serverToken {
		return response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED)
	}
	err = e.OAuth2TokenCache.Del(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	token, err := e.RedisTokenStore.GetByAccess(context.Background(), clientToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	if token == nil {
		return response.Success()
	}
	accessToken := token.GetAccess()
	refreshToken := token.GetRefresh()
	err = e.RedisTokenStore.RemoveByAccess(context.Background(), accessToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	err = e.RedisTokenStore.RemoveByRefresh(context.Background(), refreshToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.Success()
}
