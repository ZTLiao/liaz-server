package system

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/minio/minio-go/v7"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/oauth2"
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
	oauth2Config  *oauth2.Config
	ossClient     *oss.Client
	cosClient     *cos.Client
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

func SetOauth2Config(config *oauth2.Config) {
	if application.oauth2Config == nil {
		application.oauth2Config = config
	}
}

func GetOauth2Config() *oauth2.Config {
	return application.oauth2Config
}

func SetOssClient(client *oss.Client) {
	if application.ossClient == nil {
		application.ossClient = client
	}
}

func GetOssClient() *oss.Client {
	return application.ossClient
}

func SetCosClient(client *cos.Client) {
	if application.cosClient == nil {
		application.cosClient = client
	}
}

func GetCosClient() *cos.Client {
	return application.cosClient
}

var application = new(Application)
