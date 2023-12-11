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
