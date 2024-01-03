package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"strconv"
	"time"
)

type RecommendCache struct {
	redis *redis.RedisUtil
}

func NewRecommendCache(redis *redis.RedisUtil) *RecommendCache {
	return &RecommendCache{redis}
}

func (e *RecommendCache) RedisKey(position int8) string {
	return e.redis.GetKey(constant.RECOMMEND, strconv.FormatInt(int64(position), 10))
}

func (e *RecommendCache) HSet(position int8, recommendType int8, recommend *resp.RecommendResp) error {
	if recommend == nil {
		return nil
	}
	data, err := json.Marshal(recommend)
	if err != nil {
		return err
	}
	_, err = e.redis.HSet(e.RedisKey(position), strconv.FormatInt(int64(recommendType), 10), string(data))
	e.redis.Expire(e.RedisKey(position), time.Duration(constant.TIME_OF_MINUTE)*time.Second)
	return err
}

func (e *RecommendCache) HGetAll(position int8) ([]resp.RecommendResp, error) {
	data, err := e.redis.HGetAll(e.RedisKey(position))
	if err != nil {
		return nil, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	var recommends = make([]resp.RecommendResp, 0)
	for _, v := range data {
		var recommend resp.RecommendResp
		err = json.Unmarshal([]byte(v), &recommend)
		if err != nil {
			return nil, err
		}
		recommends = append(recommends, recommend)
	}
	return recommends, nil
}

func (e *RecommendCache) Del(position int8) error {
	return e.redis.Del(e.RedisKey(position))
}
