package middleware

import "github.com/gin-gonic/gin"

func SecurityHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
