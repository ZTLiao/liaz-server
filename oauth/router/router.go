package router

import (
	"basic/middleware"
	"basic/storage"
	"core/config"
	"core/constant"
	"core/logger"
	"core/response"
	"core/system"
	"core/web"
	"fmt"
	"oauth/controller"
	"oauth/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

func init() {
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		ginEngine := system.GetGinEngine()
		db := system.GetXormEngine()
		redisConfig := config.SystemConfig.Redis
		oauth2Config := config.SystemConfig.Oauth2
		manager := manage.NewDefaultManager()
		redisTokenStore := oredis.NewRedisStore(&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
			Password:     redisConfig.Password,
			DB:           redisConfig.Db,
			MinIdleConns: redisConfig.MinIdleConns,
		})
		manager.MapTokenStorage(redisTokenStore)
		clientStore := store.NewClientStore()
		clientStore.Set(oauth2Config.ClientId, &models.Client{
			ID:     oauth2Config.ClientId,
			Secret: oauth2Config.ClientSecret,
			Domain: oauth2Config.AuthServerUrl,
		})
		manager.MapClientStorage(clientStore)
		manager.SetPasswordTokenCfg(&manage.Config{
			AccessTokenExp:    time.Duration(constant.OAUTH_TOKEN_FOR_EXPIRE_TIME) * time.Second,
			IsGenerateRefresh: true,
		})
		srv := server.NewDefaultServer(manager)
		srv.SetAllowGetAccessRequest(true)
		srv.SetClientInfoHandler(server.ClientFormHandler)
		srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			logger.Error("Internal Error: %s", err.Error())
			return
		})
		srv.SetResponseErrorHandler(func(re *errors.Response) {
			logger.Error("Response Error: %s", re.Error.Error())
		})
		var userPasswordAuthorizationHander = handler.UserPasswordAuthorizationHandler{
			ClientId:  oauth2Config.ClientId,
			AccountDb: storage.NewAccountDb(db),
		}
		srv.SetPasswordAuthorizationHandler(userPasswordAuthorizationHander.Authorize)
		ginEngine.POST(constant.OAUTH_TOKEN, func(c *gin.Context) {
			srv.HandleTokenRequest(c.Writer, c.Request)
		})
		wrg.Use(middleware.SignatureHandler(config.SystemConfig.Security))
		wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		r := wrg.Group("/oauth")
		{
			new(controller.OAuthSignContoller).Router(r)
			new(controller.OAuthRefreshController).Router(r)
		}
	})
}
