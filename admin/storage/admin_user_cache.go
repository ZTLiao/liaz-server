package storage

import (
	"admin/model"
	"context"
	"core/constant"
	"core/redis"
	"encoding/json"
)

type AdminUserCache struct {
}

func (e *AdminUserCache) redisKey(accessToken string) string {
	return redis.GetKey(constant.ADMIN_USER, accessToken)
}

func (e *AdminUserCache) Set(ctx context.Context, accessToken string, adminUser *model.AdminUser) {
	if adminUser == nil {
		return
	}
	val, _ := json.Marshal(adminUser)
	redis.Set(e.redisKey(accessToken), val, constant.TIME_OF_WEEK)
}

func (e *AdminUserCache) Get(ctx context.Context, accessToken string) *model.AdminUser {
	if len(accessToken) == 0 {
		return nil
	}
	val, _ := redis.Get(e.redisKey(accessToken))
	if len(val) == 0 {
		return nil
	}
	var adminUser model.AdminUser
	json.Unmarshal([]byte(val), &adminUser)
	return &adminUser
}

func (e *AdminUserCache) Del(ctx context.Context, accessToken string) {
	redis.Del(e.redisKey(accessToken))
}
