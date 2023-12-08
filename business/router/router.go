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
		r := wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		{
			new(controller.ClientController).Router(r)
		}
	})
}
