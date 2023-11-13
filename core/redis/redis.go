package redis

import (
	"context"
	"core/application"
	"core/constant"
	"core/utils"
	"time"
)

func GetKey(suffix ...string) string {
	redisKey := utils.EMPTY
	if len(suffix) > 0 {
		redisKey = constant.PREFIX
		for _, v := range suffix {
			redisKey += utils.UNDERLINE + v
		}
	}
	return redisKey
}

func Get(key string) (string, error) {
	client := application.GetRedisClient()
	return client.Get(context.TODO(), key).Result()
}

func Set(key string, val interface{}, expire int) error {
	client := application.GetRedisClient()
	return client.Set(context.TODO(), key, val, time.Duration(expire)*time.Second).Err()
}

func Del(key string) error {
	client := application.GetRedisClient()
	return client.Del(context.TODO(), key).Err()
}

func HGet(hk, key string) (string, error) {
	client := application.GetRedisClient()
	return client.HGet(context.TODO(), hk, key).Result()
}

func HDel(hk, key string) error {
	client := application.GetRedisClient()
	return client.HDel(context.TODO(), hk, key).Err()
}

func Incr(key string) error {
	client := application.GetRedisClient()
	return client.Incr(context.TODO(), key).Err()
}

func Decr(key string) error {
	client := application.GetRedisClient()
	return client.Decr(context.TODO(), key).Err()
}

func Expire(key string, dur time.Duration) error {
	client := application.GetRedisClient()
	return client.Expire(context.TODO(), key, dur).Err()
}

func TTL(key string) (time.Duration, error) {
	client := application.GetRedisClient()
	return client.TTL(context.TODO(), key).Result()
}

func Exists(key ...string) (int64, error) {
	client := application.GetRedisClient()
	return client.Exists(context.TODO(), key...).Result()
}
