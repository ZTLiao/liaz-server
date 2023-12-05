package config

import (
	"core/middleware"
	"core/system"

	"github.com/gin-gonic/gin"
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
	routerGroup.Use(middleware.ErrorHandler()).Use(middleware.CorsHandler()).Use(middleware.RequestIdHandler()).Use(middleware.LoggerHandler())
	if env == PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	system.SetGinEngine(engine)
}
