package storage

import (
	"admin/model"
	"core/constant"
	"core/redis"
	"encoding/json"
)

type AdminUserCache struct {
	redis *redis.RedisUtil
}

func NewAdminUserCache(redis *redis.RedisUtil) *AdminUserCache {
	return &AdminUserCache{redis}
}

func (e *AdminUserCache) RedisKey(accessToken string) string {
	return e.redis.GetKey(constant.ADMIN_USER, accessToken)
}

func (e *AdminUserCache) Set(accessToken string, adminUser *model.AdminUser) error {
	if adminUser == nil {
		return nil
	}
	data, err := json.Marshal(adminUser)
	if err != nil {
		return err
	}
	e.redis.Set(e.RedisKey(accessToken), data, constant.TIME_OF_WEEK)
	return nil
}

func (e *AdminUserCache) Get(accessToken string) (*model.AdminUser, error) {
	if len(accessToken) == 0 {
		return nil, nil
	}
	data, err := e.redis.Get(e.RedisKey(accessToken))
	if err != nil {
		return nil, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	var adminUser model.AdminUser
	err = json.Unmarshal([]byte(data), &adminUser)
	if err != nil {
		return nil, err
	}
	return &adminUser, nil
}

func (e *AdminUserCache) Del(accessToken string) error {
	err := e.redis.Del(e.RedisKey(accessToken))
	if err != nil {
		return nil
	}
	return nil
}
