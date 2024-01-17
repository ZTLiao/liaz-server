package storage

import (
	"core/constant"
	"core/redis"
)

type VerifyCodeCache struct {
	redis *redis.RedisUtil
}

func NewVerifyCodeCache(redis *redis.RedisUtil) *VerifyCodeCache {
	return &VerifyCodeCache{redis}
}

func (e *VerifyCodeCache) RedisKey(username string) string {
	return e.redis.GetKey(constant.VERIFY_CODE, username)
}

func (e *VerifyCodeCache) Set(username string, verifyCode string) error {
	return e.redis.Set(e.RedisKey(username), verifyCode, constant.TIME_OF_MINUTE)
}

func (e *VerifyCodeCache) Get(username string) (string, error) {
	data, err := e.redis.Get(e.RedisKey(username))
	if err != nil {
		return "", nil
	}
	if len(data) == 0 {
		return "", nil
	}
	return data, nil
}
