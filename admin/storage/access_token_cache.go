package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type AccessTokenCache struct {
	redis *redis.RedisUtil
}

func NewAccessTokenCache(redis *redis.RedisUtil) *AccessTokenCache {
	return &AccessTokenCache{redis}
}

func (e *AccessTokenCache) RedisKey(adminId int64) string {
	return e.redis.GetKey(constant.ACCESS_TOKEN, strconv.FormatInt(adminId, 10))
}

func (e *AccessTokenCache) Set(adminId int64, accessToken string) error {
	if len(accessToken) == 0 {
		return nil
	}
	return e.redis.Set(e.RedisKey(adminId), accessToken, constant.TIME_OF_WEEK)
}

func (e *AccessTokenCache) Get(adminId int64) (string, error) {
	data, err := e.redis.Get(e.RedisKey(adminId))
	if err != nil {
		return "", nil
	}
	if len(data) == 0 {
		return "", nil
	}
	return data, nil
}

func (e *AccessTokenCache) Del(adminId int64) error {
	return e.redis.Del(e.RedisKey(adminId))
}

func (e *AccessTokenCache) TTL(adminId int64) (int64, error) {
	duration, err := e.redis.TTL(e.RedisKey(adminId))
	if err != nil {
		return 0, nil
	}
	return int64(duration.Milliseconds()), nil
}

func (e *AccessTokenCache) IsExist(adminId int64) (bool, error) {
	num, err := e.redis.Exists(e.RedisKey(adminId))
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}
