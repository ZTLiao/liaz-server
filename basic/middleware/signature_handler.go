package middleware

import (
	"core/config"
	"core/constant"
	"core/errors"
	"core/logger"
	"core/request"
	"core/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SignatureHandler(security *config.Security) gin.HandlerFunc {
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
		clientSign := c.Request.Header.Get(constant.X_SIGN)
		timestampStr := c.Request.Header.Get(constant.X_TIMESTAMP)
		var timestamp int64
		var err error
		if len(timestampStr) > 0 {
			timestamp, err = strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}
		queryParams := request.GetQueryParams(c)
		postFormParams, err := request.GetPostFormParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		var params = make(map[string][]string, 0)
		for key := range queryParams {
			params[key] = []string{fmt.Sprintf("%s", queryParams[key])}
		}
		for key := range postFormParams {
			params[key] = []string{fmt.Sprintf("%s", postFormParams[key])}
		}
		serverSign := utils.GenerateSign(params, timestamp, security.SignKey)
		if clientSign != serverSign {
			logger.Error("clientSign : %s, serverSign : %s, timestamp : %d", clientSign, serverSign, timestamp)
			c.AbortWithError(http.StatusForbidden, errors.New(http.StatusForbidden, constant.ILLEGAL_REQUEST))
		}
		c.Next()
	}
}
