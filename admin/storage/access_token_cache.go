package storage

import (
	"context"
	"core/constant"
	"core/redis"
	"strconv"
)

type AccessTokenCache struct {
}

func (e *AccessTokenCache) redisKey(adminId int64) string {
	return redis.GetKey(constant.ACCESS_TOKEN, strconv.FormatInt(adminId, 10))
}

func (e *AccessTokenCache) Set(ctx context.Context, adminId int64, accessToken string) {
	if len(accessToken) == 0 {
		return
	}
	redis.Set(e.redisKey(adminId), accessToken, constant.TIME_OF_WEEK)
}

func (e *AccessTokenCache) Get(ctx context.Context, adminId int64) string {
	data, _ := redis.Get(e.redisKey(adminId))
	if len(data) == 0 {
		return ""
	}
	return data
}

func (e *AccessTokenCache) Del(ctx context.Context, adminId int64) {
	redis.Del(e.redisKey(adminId))
}

func (e *AccessTokenCache) TTL(ctx context.Context, adminId int64) int64 {
	duration, _ := redis.TTL(e.redisKey(adminId))
	return int64(duration.Milliseconds())
}

func (e *AccessTokenCache) IsExist(ctx context.Context, adminId int64) bool {
	num, _ := redis.Exists(e.redisKey(adminId))
	return num > 0
}
