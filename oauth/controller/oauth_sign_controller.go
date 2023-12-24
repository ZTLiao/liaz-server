package controller

import (
	"basic/storage"
	"core/redis"
	"core/system"
	"core/web"
	"oauth/handler"
)

type OAuthSignContoller struct {
}

var _ web.IWebController = &OAuthSignContoller{}

func (e *OAuthSignContoller) Router(iWebRoutes web.IWebRoutes) {
	oauth2Config := system.GetOauth2Config()
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var oauthSignHander = handler.OAuthSignHandler{
		OAuth2Config:         oauth2Config,
		AccountDb:            storage.NewAccountDb(db),
		AccountLoginRecordDb: storage.NewAccountLoginRecordDb(db),
		UserDeviceDb:         storage.NewUserDeviceDb(db),
		UserDb:               storage.NewUserDb(db),
		OAuth2TokenCache:     storage.NewOAuth2TokenCache(redis),
	}
	iWebRoutes.POST("/sign/in", oauthSignHander.SignIn)
	iWebRoutes.POST("/sign/up", oauthSignHander.SignUp)
}
