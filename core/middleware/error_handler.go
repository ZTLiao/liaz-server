package middleware

import (
	"core/errors"
	"core/response"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 异常处理
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				err := fmt.Sprintf("%s", r)
				c.JSON(http.StatusOK, response.ReturnError(http.StatusInternalServerError, err))
				c.Abort()
			}
		}()
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			code := http.StatusInternalServerError
			message := err.Error()
			if apiError, ok := err.(*errors.ApiError); ok {
				if apiError.Code != 0 {
					code = apiError.Code
				}
				message = apiError.Message
			}
			c.JSON(http.StatusOK, response.ReturnError(code, message))
			c.Abort()
		}
	}
}
