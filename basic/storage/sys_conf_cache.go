package storage

import (
	"basic/model"
	"core/constant"
	"core/redis"
	"encoding/json"
)

type SysConfCache struct {
	redis *redis.RedisUtil
}

func NewSysConfCache(redis *redis.RedisUtil) *SysConfCache {
	return &SysConfCache{redis}
}

func (e *SysConfCache) RedisKey() string {
	return e.redis.GetKey(constant.SYS_CONF)
}

func (e *SysConfCache) HSet(key string, sysConf *model.SysConf) error {
	if sysConf == nil {
		return nil
	}
	val, err := json.Marshal(sysConf)
	if err != nil {
		return err
	}
	_, err = e.redis.HSet(e.RedisKey(), key, string(val))
	if err != nil {
		return nil
	}
	return nil
}

func (e *SysConfCache) HGet(key string) (*model.SysConf, error) {
	val, err := e.redis.HGet(e.RedisKey(), key)
	if err != nil {
		return nil, nil
	}
	if val == "" {
		return nil, nil
	}
	var sysConf model.SysConf
	err = json.Unmarshal([]byte(val), &sysConf)
	if err != nil {
		return nil, err
	}
	return &sysConf, nil
}

func (e *SysConfCache) Del() error {
	return e.redis.Del(e.RedisKey())
}

func (e *SysConfCache) HExists(key string) (bool, error) {
	return e.redis.HExists(e.RedisKey(), key)
}

func (e *SysConfCache) HGetAll() (map[string]model.SysConf, error) {
	var sysConfMap map[string]model.SysConf
	val, err := e.redis.HGetAll(e.RedisKey())
	if err != nil {
		return nil, nil
	}
	if len(val) == 0 {
		return nil, nil
	}
	sysConfMap = make(map[string]model.SysConf, 0)
	for k, v := range val {
		var sysConf model.SysConf
		err = json.Unmarshal([]byte(v), &sysConf)
		if err != nil {
			return nil, err
		}
		sysConfMap[k] = sysConf
	}
	return sysConfMap, nil
}
