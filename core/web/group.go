package web

import (
	"core/system"
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

type IWebController interface {
	Router(iWebRoutes IWebRoutes)
}

type WebRouterGroup struct {
	group   *gin.RouterGroup
	iRoutes gin.IRoutes
}

var _ IWebRoutes = &WebRouterGroup{}

func (e *WebRouterGroup) buildHandlersChain(handlers []WebHandlerFunc) []gin.HandlerFunc {
	var handlersChain = make([]gin.HandlerFunc, 0)
	if len(handlers) > 0 {
		for _, h := range handlers {
			handlersChain = append(handlersChain, func(ctx *gin.Context) {
				var res = h(NewWebContext(ctx))
				if len(ctx.Errors) == 0 && res != nil {
					ctx.JSON(http.StatusOK, res)
				}
			})
		}
	}
	return handlersChain
}

func (e *WebRouterGroup) Use(handlers ...gin.HandlerFunc) IWebRoutes {
	system.GetGinEngine().RouterGroup.Use(handlers...)
	return e
}

func (e *WebRouterGroup) Group(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.group.Group(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) POST(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.POST(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) GET(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.GET(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) DELETE(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.DELETE(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) PATCH(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.PATCH(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) PUT(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.PUT(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) OPTIONS(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.OPTIONS(relativePath, e.buildHandlersChain(handlers)...),
	}
}

func (e *WebRouterGroup) HEAD(relativePath string, handlers ...WebHandlerFunc) IWebRoutes {
	return &WebRouterGroup{
		group:   e.group,
		iRoutes: e.iRoutes.HEAD(relativePath, e.buildHandlersChain(handlers)...),
	}
}
