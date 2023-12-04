package router

import (
	"core/response"
	"core/web"
)

func init() {
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		{

		}
	})
}
