package middleware

import (
	"bytes"
	"core/constant"
	"core/request"
	"core/system"
	"core/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
		headers := request.GetHeaders(c)
		queryParams := request.GetQueryParams(c)
		formParams, err := request.GetPostFormParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		bodyParams, err := request.GetBodyParams(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		entry.Infof("request headers : %s, queryParams : %s, formParams : %s, bodyParams : %s", printf(headers), printf(queryParams), printf(formParams), bodyParams)
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

// 请求参数格式化日志输出
func printf(params map[string]any) string {
	var content string
	if len(params) == 0 {
		return content
	}
	for k, v := range params {
		content += k + "=" + fmt.Sprintf("%s", v) + "&"
	}
	if len(content) > 0 {
		rs := []rune(content)
		var end = len(content)
		if len(rs) < len(content) {
			end = len(rs)
		}
		content = string(rs[0 : end-1])
	}
	return content
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
