package memory

import (
	"basic/infrastructure/persistence/entity"
	"core/constant"
	"core/redis"
	"encoding/json"
)

type SysConfMemory struct {
	redis *redis.RedisUtil
}

func NewSysConfMemory(redis *redis.RedisUtil) *SysConfMemory {
	return &SysConfMemory{redis}
}

func (e *SysConfMemory) RedisKey() string {
	return e.redis.GetKey(constant.SYS_CONF)
}

func (e *SysConfMemory) HSet(key string, sysConf *entity.SysConf) error {
	if sysConf == nil {
		return nil
	}
	val, err := json.Marshal(sysConf)
	if err != nil {
		return err
	}
	_, err = e.redis.HSet(e.RedisKey(), key, string(val))
	if err != nil {
		return err
	}
	return nil
}

func (e *SysConfMemory) HGet(key string) (*entity.SysConf, error) {
	val, err := e.redis.HGet(e.RedisKey(), key)
	if err != nil {
		return nil, err
	}
	var sysConf entity.SysConf
	err = json.Unmarshal([]byte(val), &sysConf)
	if err != nil {
		return nil, err
	}
	return &sysConf, nil
}

func (e *SysConfMemory) Del() error {
	return e.redis.Del(e.RedisKey())
}

func (e *SysConfMemory) HExists(key string) (bool, error) {
	return e.redis.HExists(e.RedisKey(), key)
}

func (e *SysConfMemory) HGetAll() (map[string]entity.SysConf, error) {
	var sysConfMap map[string]entity.SysConf
	val, err := e.redis.HGetAll(e.RedisKey())
	if err != nil {
		return nil, err
	}
	sysConfMap = make(map[string]entity.SysConf, len(val))
	for k, v := range val {
		var sysConf entity.SysConf
		err = json.Unmarshal([]byte(v), &sysConf)
		if err != nil {
			return nil, err
		}
		sysConfMap[k] = sysConf
	}
	return sysConfMap, nil
}
