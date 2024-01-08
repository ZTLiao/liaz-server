package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
	"time"
)

type NovelRankCache struct {
	redis *redis.RedisUtil
}

func NewNovelRankCache(redis *redis.RedisUtil) *NovelRankCache {
	return &NovelRankCache{redis}
}

func (e *NovelRankCache) RedisKey(rankType int64, timeType int64, dateTime string) string {
	return e.redis.GetKey(constant.NOVEL_RANK, strconv.FormatInt(rankType, 10), strconv.FormatInt(timeType, 10), dateTime)
}

func (e *NovelRankCache) Incr(rankType int64, timeType int64, dateTime string, novelId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(rankType, timeType, dateTime), 1, strconv.FormatInt(novelId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *NovelRankCache) Get(rankType int64, timeType int64, dateTime string, novelId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(rankType, timeType, dateTime), strconv.FormatInt(novelId, 10))
	if err != nil {
		return 0, nil
	}
	return int64(score), nil
}

func (e *NovelRankCache) Rank(rankType int64, timeType int64, dateTime string, startIndex int64, stopIndex int64) (map[int64]int64, error) {
	res, err := e.redis.ZRevRangeWithScores(e.RedisKey(rankType, timeType, dateTime), startIndex, stopIndex)
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

func (e *NovelRankCache) Expire(rankType int64, timeType int64, dateTime string, dur time.Duration) error {
	return e.redis.Expire(e.RedisKey(rankType, timeType, dateTime), dur)
}
