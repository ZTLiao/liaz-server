package application

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Application struct {
	env           string
	name          string
	iConfigClient config_client.IConfigClient
	logger        *logrus.Logger
	ginEngine     *gin.Engine
	xormEngine    *xorm.Engine
	redisClient   *redis.Client
}

func (e *Application) SetEnv(env string) {
	if len(e.env) == 0 {
		e.env = env
	}
}

func (e *Application) GetEnv() string {
	return e.env
}

func (e *Application) SetName(name string) {
	if len(e.name) == 0 {
		e.name = name
	}
}

func (e *Application) GetName() string {
	return e.name
}

func (e *Application) SetIConfigClient(configClient config_client.IConfigClient) {
	if e.iConfigClient == nil {
		e.iConfigClient = configClient
	}
}

func (e *Application) GetIConfigClient() config_client.IConfigClient {
	return e.iConfigClient
}

func (e *Application) SetLogger(logger *logrus.Logger) {
	if e.logger == nil {
		e.logger = logger
	}
}

func (e *Application) GetLogger() *logrus.Logger {
	return e.logger
}

func (e *Application) SetGinEngine(engine *gin.Engine) {
	if e.ginEngine == nil {
		e.ginEngine = engine
	}
}

func (e *Application) GetGinEngine() *gin.Engine {
	return e.ginEngine
}

func (e *Application) SetXormEngine(engine *xorm.Engine) {
	if e.xormEngine == nil {
		e.xormEngine = engine
	}
}

func (e *Application) GetXormEngine() *xorm.Engine {
	return e.xormEngine
}

func (e *Application) SetRedisClient(client *redis.Client) {
	if e.redisClient == nil {
		e.redisClient = client
	}
}

func (e *Application) GetRedisClient() *redis.Client {
	return e.redisClient
}

var application = new(Application)

// 应用
func GetApp() *Application {
	return application
}
