package controller

import (
	"basic/storage"
	"core/config"
	coreRedis "core/redis"
	"core/system"
	"core/web"
	"fmt"
	"oauth/handler"

	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

type OAuthSignContoller struct {
}

var _ web.IWebController = &OAuthSignContoller{}

func (e *OAuthSignContoller) Router(iWebRoutes web.IWebRoutes) {
	oauth2Config := system.GetOauth2Config()
	db := system.GetXormEngine()
	var redisUtil = coreRedis.NewRedisUtil(system.GetRedisClient())
	redisConfig := config.SystemConfig.Redis
	redisTokenStore := oredis.NewRedisStore(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password:     redisConfig.Password,
		DB:           redisConfig.Db,
		MinIdleConns: redisConfig.MinIdleConns,
	})
	var oauthSignHander = handler.OAuthSignHandler{
		OAuth2Config:         oauth2Config,
		AccountDb:            storage.NewAccountDb(db),
		AccountLoginRecordDb: storage.NewAccountLoginRecordDb(db),
		UserDeviceDb:         storage.NewUserDeviceDb(db),
		UserDb:               storage.NewUserDb(db),
		OAuth2TokenCache:     storage.NewOAuth2TokenCache(redisUtil),
		RedisTokenStore:      redisTokenStore,
	}
	iWebRoutes.POST("/sign/in", oauthSignHander.SignIn)
	iWebRoutes.POST("/sign/up", oauthSignHander.SignUp)
	iWebRoutes.POST("/sign/out", oauthSignHander.SignOut)
}
