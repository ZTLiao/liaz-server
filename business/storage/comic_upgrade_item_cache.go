package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"time"
)

type ComicUpgradeItemCache struct {
	redis *redis.RedisUtil
}

func NewComicUpgradeItemCache(redis *redis.RedisUtil) *ComicUpgradeItemCache {
	return &ComicUpgradeItemCache{redis}
}

func (e *ComicUpgradeItemCache) RedisKey() string {
	return e.redis.GetKey(constant.COMIC_UPGRADE_ITEM)
}

func (e *ComicUpgradeItemCache) RPush(item resp.ComicUpgradeResp) error {
	var items = make([]string, 0)
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	items = append(items, string(data))
	_, err = e.redis.RPush(e.RedisKey(), items...)
	if err != nil {
		return err
	}
	err = e.redis.Expire(e.RedisKey(), time.Second*constant.TIME_OF_HALF_HOUR)
	if err != nil {
		return err
	}
	return nil
}

func (e *ComicUpgradeItemCache) LRange(start int64, stop int64) ([]resp.ComicUpgradeResp, error) {
	res, err := e.redis.LRange(e.RedisKey(), start, stop)
	if err != nil {
		return nil, err
	}
	var items []resp.ComicUpgradeResp
	for _, v := range res {
		var item resp.ComicUpgradeResp
		err = json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (e *ComicUpgradeItemCache) IsExist() (bool, error) {
	num, err := e.redis.Exists(e.RedisKey())
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}

func (e *ComicUpgradeItemCache) Del() error {
	return e.redis.Del(e.RedisKey())
}
