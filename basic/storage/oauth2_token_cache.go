package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type OAuth2TokenCache struct {
	redis *redis.RedisUtil
}

func NewOAuth2TokenCache(redis *redis.RedisUtil) *OAuth2TokenCache {
	return &OAuth2TokenCache{redis: redis}
}

func (e *OAuth2TokenCache) RedisKey(userId int64) string {
	return e.redis.GetKey(constant.OAUTH2_TOKEN, strconv.FormatInt(userId, 10))
}

func (e *OAuth2TokenCache) Set(userId int64, accessToken string) error {
	if len(accessToken) == 0 {
		return nil
	}
	return e.redis.Set(e.RedisKey(userId), accessToken, constant.OAUTH_TOKEN_FOR_EXPIRE_TIME)
}

func (e *OAuth2TokenCache) Get(userId int64) (string, error) {
	data, err := e.redis.Get(e.RedisKey(userId))
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", nil
	}
	return data, nil
}

func (e *OAuth2TokenCache) Del(userId int64) error {
	return e.redis.Del(e.RedisKey(userId))
}

func (e *OAuth2TokenCache) TTL(userId int64) (int64, error) {
	duration, err := e.redis.TTL(e.RedisKey(userId))
	if err != nil {
		return 0, nil
	}
	return int64(duration.Milliseconds()), nil
}

func (e *OAuth2TokenCache) IsExist(userId int64) (bool, error) {
	num, err := e.redis.Exists(e.RedisKey(userId))
	if err != nil {
		return false, err
	}
	return num > 0, nil
}
