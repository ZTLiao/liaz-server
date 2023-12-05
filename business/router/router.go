package router

import (
	"business/controller"
	"core/response"
	"core/web"
)

func init() {
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		r := wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		{
			new(controller.ClientController).Router(r)
		}
	})
}
