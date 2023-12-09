package controller

import (
	"basic/storage"
	"core/system"
	"core/web"
	"oauth/handler"
)

type OauthSignContoller struct {
}

var _ web.IWebController = &OauthSignContoller{}

func (e *OauthSignContoller) Router(iWebRoutes web.IWebRoutes) {
	oauth2Config := system.GetOauth2Config()
	db := system.GetXormEngine()
	var oauthSignHander = handler.OauthSignHandler{
		Oauth2Config: oauth2Config,
		AccountDb:    storage.NewAccountDb(db),
	}
	iWebRoutes.POST("/sign/in", oauthSignHander.SignIn)
}
