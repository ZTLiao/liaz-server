package router

import (
	"admin/controller"
	"admin/storage"
	"core/application"
	"core/config"
	"core/constant"
	"core/request"
	"core/response"
	"core/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	//设置应用名称
	application.GetApp().SetName("liaz-admin")
	//添加路由
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		wrg.Use(AdminSecurityHandler())
		wrg.Group("/").GET("/", func(wc *web.WebContext) interface{} {
			return response.Success()
		})
		r := wrg.Group("/admin")
		{
			new(controller.LoginController).Router(r)
			new(controller.AdminUserController).Router(r)
			new(controller.LogoutController).Router(r)
			new(controller.AdminMenuController).Router(r)
			new(controller.AdminRoleController).Router(r)
			new(controller.AdminRoleMenuController).Router(r)
			new(controller.AdminUserRoleController).Router(r)
		}
	})
}

func AdminSecurityHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestUri = c.Request.RequestURI
		var excludes = config.SystemConfig.Security.Excludes
		for _, v := range excludes {
			if requestUri == v {
				c.Next()
				return
			}
		}
		var accessToken = c.Request.Header.Get(constant.AUTHORIZATION)
		if len(accessToken) == 0 {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		var adminUser = new(storage.AdminUserCache).Get(accessToken)
		if adminUser == nil {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		headers := c.Request.Header
		formParams, _ := request.GetPostFormParams(c)
		queryParams := request.GetQueryParams(c)
		bodyParams := request.GetBodyParams(c)
		new(storage.AdminLogDb).AddLog(adminUser.AdminId, c.Request.RequestURI, headers, queryParams, formParams, bodyParams)
		c.Next()
	}
}
