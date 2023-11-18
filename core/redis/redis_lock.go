package redis

import (
	"core/constant"
	"core/system"
)

type RedisLock struct {
	redis    *RedisUtil
	unqiueId string
	lockVal  string
}

func NewRedisLock(unqiueId string) *RedisLock {
	redis := NewRedisUtil(system.GetRedisClient())
	return &RedisLock{
		redis:    redis,
		unqiueId: unqiueId,
	}
}

func (e *RedisLock) redisKey() string {
	return e.redis.GetKey(constant.REDIS_LOCK, e.unqiueId)
}

func (e *RedisLock) Lock() bool {
	e.lockVal = e.redis.Lock(e.redisKey(), constant.TIME_OF_MINUTE, constant.TIME_OF_MINUTE)
	return len(e.lockVal) > 0
}

func (e *RedisLock) Unlock() {
	e.redis.Unlock(e.redisKey(), e.lockVal)
}
