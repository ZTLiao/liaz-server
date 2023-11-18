package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/minio/minio-go/v7"
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
	minioClient   *minio.Client
}

func SetEnv(env string) {
	if len(application.env) == 0 {
		application.env = env
	}
}

func GetEnv() string {
	return application.env
}

func SetName(name string) {
	if len(application.name) == 0 {
		application.name = name
	}
}

func GetName() string {
	return application.name
}

func SetIConfigClient(configClient config_client.IConfigClient) {
	if application.iConfigClient == nil {
		application.iConfigClient = configClient
	}
}

func GetIConfigClient() config_client.IConfigClient {
	return application.iConfigClient
}

func SetLogger(logger *logrus.Logger) {
	if application.logger == nil {
		application.logger = logger
	}
}

func GetLogger() *logrus.Logger {
	return application.logger
}

func SetGinEngine(engine *gin.Engine) {
	if application.ginEngine == nil {
		application.ginEngine = engine
	}
}

func GetGinEngine() *gin.Engine {
	return application.ginEngine
}

func SetXormEngine(engine *xorm.Engine) {
	if application.xormEngine == nil {
		application.xormEngine = engine
	}
}

func GetXormEngine() *xorm.Engine {
	return application.xormEngine
}

func SetRedisClient(client *redis.Client) {
	if application.redisClient == nil {
		application.redisClient = client
	}
}

func GetRedisClient() *redis.Client {
	return application.redisClient
}

func SetMinioClient(client *minio.Client) {
	if application.minioClient == nil {
		application.minioClient = client
	}
}

func GetMinioClient() *minio.Client {
	return application.minioClient
}

var application = new(Application)
