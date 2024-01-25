package router

import (
	"admin/controller"
	"admin/storage"
	"core/config"
	"core/constant"
	"core/redis"
	"core/request"
	"core/response"
	"core/system"
	"core/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func init() {
	//添加路由
	web.AddRouter(func(wrg *web.WebRouterGroup) {
		db := system.GetXormEngine()
		redis := redis.NewRedisUtil(system.GetRedisClient())
		wrg.Use(AdminSecurityHandler(config.SystemConfig.Security, db, redis))
		var success = func(wc *web.WebContext) interface{} {
			return response.Success()
		}
		wrg.Group("/").GET("/", success).HEAD("/", success)
		r := wrg.Group("/admin")
		{
			new(controller.AdminLoginController).Router(r)
			new(controller.AdminLogoutController).Router(r)
			new(controller.AdminUserController).Router(r)
			new(controller.AdminMenuController).Router(r)
			new(controller.AdminRoleController).Router(r)
			new(controller.AdminRoleMenuController).Router(r)
			new(controller.AdminUserRoleController).Router(r)
			new(controller.AdminFileController).Router(r)
			new(controller.AdminSysConfController).Router(r)
			new(controller.AdminRecommendController).Router(r)
			new(controller.AdminRecommendItemController).Router(r)
			new(controller.AdminCategoryController).Router(r)
			new(controller.AdminAuthorController).Router(r)
		}
	})
}

func AdminSecurityHandler(security *config.Security, db *xorm.Engine, redis *redis.RedisUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !security.Encrypt {
			c.Next()
			return
		}
		requestUri := c.Request.RequestURI
		excludes := security.Excludes
		for _, v := range excludes {
			if requestUri == v {
				c.Next()
				return
			}
		}
		accessToken := c.Request.Header.Get(constant.AUTHORIZATION)
		if len(accessToken) == 0 {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		adminUser, err := storage.NewAdminUserCache(redis).Get(accessToken)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		if adminUser == nil {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		headers := request.GetHeaders(c)
		formParams, err := request.GetPostFormParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		queryParams := request.GetQueryParams(c)
		bodyParams, err := request.GetBodyParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		storage.NewAdminLogDb(db).AddLog(adminUser.AdminId, c.Request.RequestURI, headers, queryParams, formParams, bodyParams)
		c.Next()
	}
}
