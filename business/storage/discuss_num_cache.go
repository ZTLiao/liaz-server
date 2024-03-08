package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type DiscussNumCache struct {
	redis *redis.RedisUtil
}

func NewDiscussNumCache(redis *redis.RedisUtil) *DiscussNumCache {
	return &DiscussNumCache{redis}
}

func (e *DiscussNumCache) RedisKey() string {
	return e.redis.GetKey(constant.DISCUSS_NUM)
}

func (e *DiscussNumCache) Incr(discussId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(), 1, strconv.FormatInt(discussId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *DiscussNumCache) Decr(discussId int64) error {
	_, err := e.redis.ZIncrBy(e.RedisKey(), -1, strconv.FormatInt(discussId, 10))
	if err != nil {
		return err
	}
	return nil
}

func (e *DiscussNumCache) Get(discussId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(), strconv.FormatInt(discussId, 10))
	if err != nil {
		return 0, nil
	}
	return int64(score), nil
}

func (e *DiscussNumCache) Rank(startIndex int64, stopIndex int64) (map[int64]int64, error) {
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
