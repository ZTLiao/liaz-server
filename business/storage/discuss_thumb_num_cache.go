package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type DiscussThumbNumCache struct {
	redis *redis.RedisUtil
}

func NewDiscussThumbNumCache(redis *redis.RedisUtil) *DiscussThumbNumCache {
	return &DiscussThumbNumCache{redis}
}

func (e *DiscussThumbNumCache) RedisKey() string {
	return e.redis.GetKey(constant.DISCUSS_THUMB_NUM)
}

func (e *DiscussThumbNumCache) Incr(discussId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(), 1, strconv.FormatInt(discussId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *DiscussThumbNumCache) Decr(discussId int64) error {
	_, err := e.redis.ZIncrBy(e.RedisKey(), -1, strconv.FormatInt(discussId, 10))
	if err != nil {
		return err
	}
	return nil
}

func (e *DiscussThumbNumCache) Get(discussId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(), strconv.FormatInt(discussId, 10))
	if err != nil {
		return 0, nil
	}
	return int64(score), nil
}

func (e *DiscussThumbNumCache) Rank(startIndex int64, stopIndex int64) (map[int64]int64, error) {
	res, err := e.redis.ZRevRangeWithScores(e.RedisKey(), startIndex, stopIndex)
	if err != nil {
		return nil, nil
	}
	var data = make(map[int64]int64, 0)
	if len(res) == 0 {
		return data, nil
	}
	for _, v := range res {
		member, err := strconv.ParseInt(v.Member.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		data[member] = int64(v.Score)
	}
	return data, nil
}
