package redis

import (
	"context"
	"core/constant"
	"core/logger"
	"core/utils"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUtil struct {
	db *redis.Client
}

func NewRedisUtil(db *redis.Client) *RedisUtil {
	return &RedisUtil{db}
}

func (e *RedisUtil) GetKey(suffix ...string) string {
	var redisKey string
	if len(suffix) > 0 {
		redisKey = constant.PREFIX
		for _, v := range suffix {
			redisKey += utils.UNDERLINE + v
		}
	}
	return redisKey
}

func (e *RedisUtil) Lock(key string, waitTime int, timeout int) string {
	var lockValue int64
	start := time.Now().UnixMilli()
	defer func() {
		e.Expire(key, time.Duration(timeout))
		end := time.Now().UnixMilli()
		logger.Info("redis lock key : %s, waitTime : %d, timeout : %d, runTime : %d, lockValue : %d", key, waitTime, timeout, (end - start), lockValue)
	}()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("redis lock err is %s", err)
		}
	}()
	retryTime := waitTime
	for retryTime > 0 {
		lockValue = time.Now().UnixNano()
		if ok, err := e.SetNX(key, strconv.FormatInt(lockValue, 10), time.Duration(constant.TIME_OF_MINUTE)); err == nil && ok {
			return strconv.FormatInt(lockValue, 10)
		} else {
			logger.Error(err.Error())
		}
		time.Sleep(time.Second)
	}
	return ""
}

func (e *RedisUtil) Unlock(key string, lockVal string) error {
	value, err := e.Get(key)
	if err != nil {
		return err
	}
	if len(value) > 0 && value == lockVal {
		e.Del(key)
		logger.Info("redis unlock key : %s, lockVal : %s", key, lockVal)
	}
	return nil
}

func (e *RedisUtil) Get(key string) (string, error) {
	count, err := e.db.Exists(context.TODO(), key).Result()
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "", nil
	}
	res, err := e.db.Get(context.TODO(), key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) Set(key string, val interface{}, expire int) error {
	err := e.db.Set(context.TODO(), key, val, time.Duration(expire)*time.Second).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) Del(key string) error {
	err := e.db.Del(context.TODO(), key).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) HSet(hk string, value ...string) (int64, error) {
	res, err := e.db.HSet(context.TODO(), hk, value).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) HGet(hk, key string) (string, error) {
	count, err := e.db.Exists(context.TODO(), hk).Result()
	if err != nil {
		return "", err
	}
	if count == 0 {
		return "", nil
	}
	res, err := e.db.HGet(context.TODO(), hk, key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) HDel(hk, key string) error {
	err := e.db.HDel(context.TODO(), hk, key).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) HGetAll(hk string) (map[string]string, error) {
	count, err := e.db.Exists(context.TODO(), hk).Result()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	res, err := e.db.HGetAll(context.TODO(), hk).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) HExists(hk, key string) (bool, error) {
	res, err := e.db.HExists(context.TODO(), hk, key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) Incr(key string) error {
	err := e.db.Incr(context.TODO(), key).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) Decr(key string) error {
	err := e.db.Decr(context.TODO(), key).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) Expire(key string, dur time.Duration) error {
	err := e.db.Expire(context.TODO(), key, dur).Err()
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (e *RedisUtil) TTL(key string) (time.Duration, error) {
	res, err := e.db.TTL(context.TODO(), key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) Exists(key ...string) (int64, error) {
	res, err := e.db.Exists(context.TODO(), key...).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) SetNX(key string, value interface{}, dur time.Duration) (bool, error) {
	res, err := e.db.SetNX(context.TODO(), key, value, dur).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}
