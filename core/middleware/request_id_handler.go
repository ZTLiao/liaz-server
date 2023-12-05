package middleware

import (
	"core/constant"
	"core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
)

// 请求ID拦截器
func RequestIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		X_REQUEST_ID := constant.X_REQUEST_ID
		requestId := c.Request.Header.Get(X_REQUEST_ID)
		if requestId == utils.EMPTY {
			uuid, err := uuid.NewV4()
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			requestId = uuid.String()
		}
		c.Set(X_REQUEST_ID, requestId)
		c.Request.Header.Set(X_REQUEST_ID, requestId)
		c.Next()
	}
}
