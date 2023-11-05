package config

import (
	"core/application"
	"core/constant"
	"core/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Port int `yaml:"port"`
}

func (server *Server) Init() {
	var env = application.GetApp().GetEnv()
	var engine = gin.New()
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	if env == PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	engine.Use(RequestIdHandler())
	engine.Use(LoggerHandler())
	application.GetApp().SetEngine(engine)
}

// 请求ID拦截器
func RequestIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		X_REQUEST_ID := constant.X_REQUEST_ID
		requestId := c.Request.Header.Get(X_REQUEST_ID)
		if requestId == utils.EMPTY {
			uuid, _ := uuid.NewV4()
			requestId = uuid.String()
		}
		c.Set(X_REQUEST_ID, requestId)
		c.Writer.Header().Set(X_REQUEST_ID, requestId)
		c.Next()
	}
}

// 日志拦截器
func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := application.GetApp().GetLogger()
		startTime := time.Now()
		c.Next()
		spendTime := time.Since(startTime).Milliseconds()
		//API调用耗时
		ST := fmt.Sprintf("%d ms", spendTime)
		//状态码
		statusCode := c.Writer.Status()
		//请求客户端的IP
		clientIP := c.ClientIP()
		//请求方法
		method := c.Request.Method
		//请求URL
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"Status":    statusCode,
			"SpendTime": ST,
			"IP":        clientIP,
			"Method":    method,
			"Path":      path,
		})
		//Errors保存了使用当前context的所有中间件/handler所产生的全部错误信息。
		if len(c.Errors) > 0 {
			logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		//根据状态码决定打印log的等级
		if statusCode >= http.StatusInternalServerError {
			entry.Error()
		} else if statusCode >= http.StatusBadRequest {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
