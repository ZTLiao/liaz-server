package handler

import (
	"basic/storage"
	"context"
	"core/response"
	"core/types"
	"core/web"
	"oauth/resp"

	"golang.org/x/oauth2"
)

type OauthSignHandler struct {
	Oauth2Config *oauth2.Config
	AccountDb    *storage.AccountDb
}

func (e *OauthSignHandler) SignIn(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	password := wc.PostForm("password")
	token, err := e.Oauth2Config.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		if e, ok := err.(*oauth2.RetrieveError); ok {
			wc.Error("errorCode : %s, errorDescription : %s", e.ErrorCode, e.ErrorDescription)
			return response.Fail(e.ErrorDescription)
		}
		wc.AbortWithError(err)
	}
	account, err := e.AccountDb.SignIn(username, password)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := account.UserId
	return response.ReturnOK(&resp.CustomOAuth2Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       types.Time(token.Expiry),
		UserId:       userId,
	})
}