package config

import (
	"context"
	"core/system"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Password       string `yaml:"password"`
	Db             int    `yaml:"db"`
	MinIdleConns   int    `yaml:"minIdleConns"`
	MaxIdleConns   int    `yaml:"maxIdleConns"`
	MaxActiveConns int    `yaml:"maxActiveConns"`
}

func (e *Redis) Init() {
	if e == nil {
		return
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:           fmt.Sprintf("%s:%s", e.Host, e.Port),
		Password:       e.Password,
		DB:             e.Db,
		MinIdleConns:   e.MinIdleConns,
		MaxIdleConns:   e.MaxIdleConns,
		MaxActiveConns: e.MaxActiveConns,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	system.SetRedisClient(redisClient)
}
