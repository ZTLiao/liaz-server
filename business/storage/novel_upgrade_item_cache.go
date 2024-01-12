package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"time"
)

type NovelUpgradeItemCache struct {
	redis *redis.RedisUtil
}

func NewNovelUpgradeItemCache(redis *redis.RedisUtil) *NovelUpgradeItemCache {
	return &NovelUpgradeItemCache{redis}
}

func (e *NovelUpgradeItemCache) RedisKey() string {
	return e.redis.GetKey(constant.NOVEL_UPGRADE_ITEM)
}

func (e *NovelUpgradeItemCache) LPush(item resp.NovelUpgradeResp) error {
	var items = make([]string, 0)
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	items = append(items, string(data))
	_, err = e.redis.LPush(e.RedisKey(), items...)
	if err != nil {
		return err
	}
	err = e.redis.Expire(e.RedisKey(), time.Second*constant.TIME_OF_HOUR)
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelUpgradeItemCache) LRange(start int64, stop int64) ([]resp.NovelUpgradeResp, error) {
	res, err := e.redis.LRange(e.RedisKey(), start, stop)
	if err != nil {
		return nil, err
	}
	var items []resp.NovelUpgradeResp
	for _, v := range res {
		var item resp.NovelUpgradeResp
		err = json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (e *NovelUpgradeItemCache) IsExist() (bool, error) {
	num, err := e.redis.Exists(e.RedisKey())
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}
