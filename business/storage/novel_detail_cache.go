package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"strconv"
)

type NovelDetailCache struct {
	redis *redis.RedisUtil
}

func NewNovelDetailCache(redis *redis.RedisUtil) *NovelDetailCache {
	return &NovelDetailCache{redis}
}

func (e *NovelDetailCache) RedisKey(novelId int64) string {
	return e.redis.GetKey(constant.NOVEL_DETAIL, strconv.FormatInt(novelId, 10))
}

func (e *NovelDetailCache) Set(novelId int64, novelDetail *resp.NovelDetailResp) error {
	if novelDetail == nil {
		return nil
	}
	data, err := json.Marshal(novelDetail)
	if err != nil {
		return err
	}
	return e.redis.Set(e.RedisKey(novelId), data, constant.TIME_OF_WEEK)
}

func (e *NovelDetailCache) Get(novelId int64) (*resp.NovelDetailResp, error) {
	data, err := e.redis.Get(e.RedisKey(novelId))
	if err != nil {
		return nil, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	var novelDetail resp.NovelDetailResp
	err = json.Unmarshal([]byte(data), &novelDetail)
	if err != nil {
		return nil, err
	}
	return &novelDetail, nil
}

func (e *NovelDetailCache) Del(novelId int64) error {
	return e.redis.Del(e.RedisKey(novelId))
}

func (e *NovelDetailCache) TTL(novelId int64) (int64, error) {
	duration, err := e.redis.TTL(e.RedisKey(novelId))
	if err != nil {
		return 0, nil
	}
	return int64(duration.Milliseconds()), nil
}

func (e *NovelDetailCache) IsExist(novelId int64) (bool, error) {
	num, err := e.redis.Exists(e.RedisKey(novelId))
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}
