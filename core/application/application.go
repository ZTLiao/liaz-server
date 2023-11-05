package application

import (
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/sirupsen/logrus"
)

type Application struct {
	env           string
	name          string
	iConfigClient config_client.IConfigClient
	logger        *logrus.Logger
	engine        *gin.Engine
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

func (e *Application) SetEngine(engine *gin.Engine) {
	if e.engine == nil {
		e.engine = engine
	}
}

func (e *Application) GetEngine() *gin.Engine {
	return e.engine
}

var application = new(Application)

// 应用
func GetApp() *Application {
	return application
}
