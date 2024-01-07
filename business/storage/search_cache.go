package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type SearchCache struct {
	redis *redis.RedisUtil
}

func NewSearchCache(redis *redis.RedisUtil) *SearchCache {
	return &SearchCache{redis}
}

func (e *SearchCache) RedisKey() string {
	return e.redis.GetKey(constant.SEARCH)
}

func (e *SearchCache) Incr(assetId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(), 1, strconv.FormatInt(assetId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *SearchCache) Get(assetId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(), strconv.FormatInt(assetId, 10))
	if err != nil {
		return 0, nil
	}
	return int64(score), nil
}

func (e *SearchCache) Rank(startIndex int64, stopIndex int64) (map[int64]int64, error) {
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
