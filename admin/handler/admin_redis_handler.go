package handler

import (
	"core/redis"
	"core/response"
	"core/web"
)

type AdminRedisHandler struct {
	RedisUtil *redis.RedisUtil
}

func (e *AdminRedisHandler) Delete(wc *web.WebContext) interface{} {
	redisKey := wc.PostForm("redisKey")
	err := e.RedisUtil.Del(redisKey)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.Success()
}
