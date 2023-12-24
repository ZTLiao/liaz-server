package handler

import (
	"basic/storage"
	"context"
	"core/constant"
	"core/response"
	"core/types"
	"core/web"
	"net/http"
	"oauth/resp"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

type OAuthRefreshHandler struct {
	OAuth2Config     *oauth2.Config
	OAuth2TokenCache *storage.OAuth2TokenCache
}

func (e *OAuthRefreshHandler) RefreshToken(wc *web.WebContext) interface{} {
	refreshToken := wc.PostForm("token")
	userIdStr := wc.GetHeader(constant.X_USER_ID)
	if len(userIdStr) == 0 {
		return response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED)
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	accessToken, err := e.OAuth2TokenCache.Get(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	var oauth2Token = new(oauth2.Token)
	oauth2Token.AccessToken = accessToken
	oauth2Token.RefreshToken = refreshToken
	oauth2Token.Expiry = time.Now()
	token, err := e.OAuth2Config.TokenSource(context.Background(), oauth2Token).Token()
	if err != nil {
		if e, ok := err.(*oauth2.RetrieveError); ok {
			wc.Error("errorCode : %s, errorDescription : %s", e.ErrorCode, e.ErrorDescription)
			return response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED)
		}
		wc.AbortWithError(err)
	}
	e.OAuth2TokenCache.Set(userId, token.AccessToken)
	return response.ReturnOK(&resp.CustomOAuth2TokenResp{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       types.Time(token.Expiry),
		UserId:       userId,
	})
}
