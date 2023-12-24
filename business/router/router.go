package router

import (
	"basic/middleware"
	"business/controller"
	"core/config"
	"core/response"
	"core/web"
)

func init() {
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		wrg.Use(middleware.SignatureHandler(config.SystemConfig.Security))
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
		}
	})
}
