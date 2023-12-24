package middleware

import (
	"basic/storage"
	"core/config"
	"core/constant"
	"core/response"
	"core/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SecurityHandler(security *config.Security, oauth2TokenCache *storage.OAuth2TokenCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUri := c.Request.RequestURI
		authorizes := security.Authorizes
		var isExist bool = false
		for _, v := range authorizes {
			if requestUri == v {
				isExist = true
				break
			}
		}
		if !isExist {
			c.Next()
			return
		}
		clientToken := c.Request.Header.Get(constant.AUTHORIZATION)
		if len(clientToken) == 0 {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		tokenArray := strings.Split(clientToken, utils.SPACE)
		tokenType := tokenArray[0]
		if tokenType != constant.TOKEN_TYPE {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		clientToken = tokenArray[1]
		userIdStr := c.Request.Header.Get(constant.X_USER_ID)
		if len(userIdStr) == 0 {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		serverToken, err := oauth2TokenCache.Get(userId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if clientToken != serverToken {
			c.JSON(http.StatusUnauthorized, response.ReturnError(http.StatusUnauthorized, constant.UNAUTHORIZED))
			c.Abort()
			return
		}
		c.Next()
	}
}
