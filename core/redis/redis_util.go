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

func (e *RedisUtil) ZAdd(key string, member string, score float64) (int64, error) {
	res, err := e.db.ZAdd(context.TODO(), key, redis.Z{
		Score:  score,
		Member: member,
	}).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) ZIncrBy(key string, increment float64, member string) (float64, error) {
	res, err := e.db.ZIncrBy(context.TODO(), key, increment, member).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) ZScore(key string, member string) (float64, error) {
	res, err := e.db.ZScore(context.TODO(), key, member).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) ZRangeWithScores(key string, start int64, stop int64) ([]redis.Z, error) {
	res, err := e.db.ZRangeWithScores(context.TODO(), key, start, stop).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) ZRevRangeWithScores(key string, start int64, stop int64) ([]redis.Z, error) {
	res, err := e.db.ZRevRangeWithScores(context.TODO(), key, start, stop).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) LPush(key string, values ...string) (int64, error) {
	res, err := e.db.LPush(context.TODO(), key, values).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) RPush(key string, values ...string) (int64, error) {
	res, err := e.db.RPush(context.TODO(), key, values).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) LPop(key string) (string, error) {
	res, err := e.db.LPop(context.TODO(), key).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}

func (e *RedisUtil) LRange(key string, start int64, stop int64) ([]string, error) {
	res, err := e.db.LRange(context.TODO(), key, start, stop).Result()
	if err != nil {
		logger.Error(err.Error())
	}
	return res, err
}
