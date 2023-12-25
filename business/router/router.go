package router

import (
	"basic/middleware"
	"basic/storage"
	"business/controller"
	"core/config"
	"core/redis"
	"core/response"
	"core/system"
	"core/web"
)

func init() {
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		security := config.SystemConfig.Security
		var redis = redis.NewRedisUtil(system.GetRedisClient())
		wrg.Use(middleware.SignatureHandler(security))
		wrg.Use(middleware.SecurityHandler(security, storage.NewOAuth2TokenCache(redis)))
		wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		r := wrg.Group("/api")
		{
			new(controller.ClientController).Router(r)
			new(controller.RecommendController).Router(r)
			new(controller.CategoryController).Router(r)
			new(controller.ComicController).Router(r)
			new(controller.CategorySearchController).Router(r)
			new(controller.ComicChapterController).Router(r)
			new(controller.FileController).Router(r)
			new(controller.UserController).Router(r)
			new(controller.NovelController).Router(r)
		}
	})
}
