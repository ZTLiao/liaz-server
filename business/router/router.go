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
			new(controller.NovelChapterController).Router(r)
			new(controller.BrowseController).Router(r)
			new(controller.ComicSubscribeController).Router(r)
			new(controller.NovelSubscribeController).Router(r)
			new(controller.BookshelfController).Router(r)
			new(controller.SearchController).Router(r)
			new(controller.RankController).Router(r)
			new(controller.VerifyCodeController).Router(r)
			new(controller.AccountController).Router(r)
			new(controller.CrashRecordController).Router(r)
		}
	})
}
