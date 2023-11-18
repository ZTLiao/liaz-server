package config

import (
	"bytes"
	"core/constant"
	"core/errors"
	"core/request"
	"core/response"
	"core/system"
	"core/utils"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Port int `yaml:"port"`
}

func (e *Server) Init() {
	if e == nil {
		return
	}
	env := system.GetEnv()
	engine := gin.New()
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	routerGroup := engine.RouterGroup
	routerGroup.Use(ErrorHandler()).Use(CorsHandler()).Use(RequestIdHandler()).Use(LoggerHandler())
	if env == PROD {
		gin.SetMode(gin.ReleaseMode)
	} else {
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	}
	system.SetGinEngine(engine)
}

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

// 跨域
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

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

// 日志拦截器
func LoggerHandler() gin.HandlerFunc {
	logger := system.GetLogger()
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(constant.X_REQUEST_ID)
		//请求客户端的IP
		clientIP := c.ClientIP()
		//请求方法
		method := c.Request.Method
		//请求URL
		path := c.Request.RequestURI
		startTime := time.Now()
		//日志体
		entry := logger.WithFields(logrus.Fields{
			"requestId": requestId,
			"clientIp":  clientIP,
			"method":    method,
			"path":      path,
		})
		//封装响应体
		writerWrapper := &ResponseWriterWrapper{Body: bytes.NewBufferString(utils.EMPTY), ResponseWriter: c.Writer}
		c.Writer = writerWrapper
		queryParams := request.GetQueryParams(c)
		formParams, err := request.GetPostFormParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		bodyParams, err := request.GetBodyParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		entry.Infof("request headers : %s, queryParams : %s, formParams : %s, bodyParams : %s", c.Request.Header, queryParams, formParams, bodyParams)
		c.Next()
		spendTime := time.Since(startTime).Milliseconds()
		//API调用耗时
		ST := fmt.Sprintf("%d ms", spendTime)
		//状态码
		statusCode := c.Writer.Status()
		//Errors保存了使用当前context的所有中间件/handler所产生的全部错误信息。
		if len(c.Errors) > 0 {
			logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		entry.Infof("response status : %s, spendTime : %s, result : %s", strconv.FormatInt(int64(statusCode), 10), ST, writerWrapper.Body.String())
	}
}

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
