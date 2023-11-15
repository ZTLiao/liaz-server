package web

import (
	"core/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebHandlerFunc func(*WebContext) interface{}

type IWebRoutes interface {
	GET(string, ...WebHandlerFunc) IWebRoutes
	POST(string, ...WebHandlerFunc) IWebRoutes
	DELETE(string, ...WebHandlerFunc) IWebRoutes
	PATCH(string, ...WebHandlerFunc) IWebRoutes
	PUT(string, ...WebHandlerFunc) IWebRoutes
	OPTIONS(string, ...WebHandlerFunc) IWebRoutes
	HEAD(string, ...WebHandlerFunc) IWebRoutes
}

type WebRouterGroup struct {
	group   *gin.RouterGroup
	iRoutes gin.IRoutes
}

var _ IWebRoutes = &WebRouterGroup{}

func (e *WebRouterGroup) buildHandlersChain(handlers []WebHandlerFunc) []gin.HandlerFunc {
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

func (e *WebRouterGroup) Use(handlers ...gin.HandlerFunc) IWebRoutes {
	var engine = application.GetGinEngine()
	engine.RouterGroup.Use(handlers...)
	return e
}

func (e *WebRouterGroup) Group(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.group.Group(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) POST(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.POST(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) GET(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.GET(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) DELETE(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.DELETE(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) PATCH(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.PATCH(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) PUT(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.PUT(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) OPTIONS(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.OPTIONS(relativePath, handlersChain...)
	return wrg
}

func (e *WebRouterGroup) HEAD(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	var handlersChain []gin.HandlerFunc = e.buildHandlersChain(handlers)
	var wrg = new(WebRouterGroup)
	wrg.group = e.group
	wrg.iRoutes = e.iRoutes.HEAD(relativePath, handlersChain...)
	return wrg
}
