package controller

import (
	"basic/storage"
	"core/redis"
	"core/system"
	"core/web"
	"oauth/handler"
)

type OAuthRefreshController struct {
}

var _ web.IWebController = &OAuthRefreshController{}

func (e *OAuthRefreshController) Router(iWebRoutes web.IWebRoutes) {
	oauth2Config := system.GetOauth2Config()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var oauthRefreshHandler = &handler.OAuthRefreshHandler{
		OAuth2Config:     oauth2Config,
		OAuth2TokenCache: storage.NewOAuth2TokenCache(redis),
	}
	iWebRoutes.POST("/refresh/token", oauthRefreshHandler.RefreshToken)
}
