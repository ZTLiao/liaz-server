package router

import (
	"core/application"
	"core/response"
	"core/web"
	"fmt"
)

func init() {
	//设置应用名称
	application.GetApp().SetName("liaz-admin")
}

func init() {
	//添加路由
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		r := wrg.Group("/admin")
		{
			r.GET("/", func(wc *web.WebContext) interface{} {
				fmt.Println("test...")
				if true {
					panic("error ....")
					//wc.Context.Error(errors.New(503, "服务器异常"))
					//return response.Fail()
				}
				return response.Success()
			})
		}
	})
}
