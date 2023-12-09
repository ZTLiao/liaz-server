package controller

import (
	"core/system"
	"core/web"
	"oauth/handler"
)

type OauthSignContoller struct {
}

var _ web.IWebController = &OauthSignContoller{}

func (e *OauthSignContoller) Router(iWebRoutes web.IWebRoutes) {
	oauth2Config := system.GetOauth2Config()
	var oauthSignHander = handler.OauthSignHandler{
		Oauth2Config: oauth2Config,
	}
	iWebRoutes.POST("/sign/in", oauthSignHander.SignIn)
}
