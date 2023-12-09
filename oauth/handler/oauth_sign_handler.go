package handler

import (
	"context"
	"core/response"
	"core/web"

	"golang.org/x/oauth2"
)

type OauthSignHandler struct {
	Oauth2Config *oauth2.Config
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
	return response.ReturnOK(token)
}
