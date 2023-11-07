package web

import (
	"core/application"
	"core/constant"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	routers = make([]func(*WebRouterGroup), 0)
)

type WebContext struct {
	Context *gin.Context
}

func (e *WebContext) Trace(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Tracef(format, args...)
}

func (e *WebContext) Debug(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Debugf(format, args...)
}

func (e *WebContext) Info(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Infof(format, args...)
}

func (e *WebContext) Warn(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Warnf(format, args...)
}

func (e *WebContext) Error(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Errorf(format, args...)
}

func (ctx *WebContext) Fatal(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": ctx.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Fatalf(format, args...)
}

func (e *WebContext) Panic(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Panicf(format, args...)
}

type WebHandleFunc func(*WebContext) interface{}

type IWebRoutes interface {
	GET(string, ...WebHandleFunc) IWebRoutes
	POST(string, ...WebHandleFunc) IWebRoutes
	DELETE(string, ...WebHandleFunc) IWebRoutes
	PATCH(string, ...WebHandleFunc) IWebRoutes
	PUT(string, ...WebHandleFunc) IWebRoutes
	OPTIONS(string, ...WebHandleFunc) IWebRoutes
	HEAD(string, ...WebHandleFunc) IWebRoutes
}

type WebRouterGroup struct {
	group   *gin.RouterGroup
	iRoutes gin.IRoutes
}

func (e *WebRouterGroup) buildHandlersChain(handlers []WebHandleFunc) []gin.HandlerFunc {
	var handlersChain []gin.HandlerFunc
	handlersChain = make([]gin.HandlerFunc, 0)
	if len(handlers) > 0 {
		for _, h := range handlers {
			handlersChain = append(handlersChain, func(ctx *gin.Context) {
				var res = h(&WebContext{
					Context: ctx,
				})
				if len(ctx.Errors) == 0 && res != nil {
					ctx.JSON(http.StatusOK, res)
				}
			})
		}
	}
	return handlersChain
}

func (e *WebRouterGroup) Group(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.group.Group(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) POST(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.POST(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) GET(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.GET(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) DELETE(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.DELETE(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) PATCH(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.PATCH(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) PUT(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.PUT(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) OPTIONS(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.OPTIONS(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) HEAD(relativePath string, handlers ...WebHandleFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.HEAD(relativePath, handlersChain...)
	return wrg
}

// 添加路由
func AddRouter(handlers ...func(*WebRouterGroup)) {
	routers = append(routers, handlers...)
}

// 初始化路由
func InitRouter(engine *gin.Engine) {
	var wrg = new(WebRouterGroup)
	wrg.group = &engine.RouterGroup
	for _, h := range routers {
		h(wrg)
	}
}
