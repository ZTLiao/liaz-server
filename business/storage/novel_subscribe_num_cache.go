package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type NovelSubscribeNumCache struct {
	redis *redis.RedisUtil
}

func NewNovelSubscribeNumCache(redis *redis.RedisUtil) *NovelSubscribeNumCache {
	return &NovelSubscribeNumCache{redis}
}

func (e *NovelSubscribeNumCache) RedisKey() string {
	return e.redis.GetKey(constant.NOVEL_SUBSCRIBE_NUM)
}

func (e *NovelSubscribeNumCache) Incr(novelId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(), 1, strconv.FormatInt(novelId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *NovelSubscribeNumCache) Decr(novelId int64) error {
	_, err := e.redis.ZIncrBy(e.RedisKey(), -1, strconv.FormatInt(novelId, 10))
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelSubscribeNumCache) Get(novelId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(), strconv.FormatInt(novelId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *NovelSubscribeNumCache) Rank(startIndex int64, stopIndex int64) ([]map[int64]int64, error) {
	res, err := e.redis.ZRevRangeWithScores(e.RedisKey(), startIndex, stopIndex)
	if err != nil {
		return nil, err
	}
	var data = make([]map[int64]int64, 0)
	if len(res) == 0 {
		return data, nil
	}
	for _, v := range res {
		member, err := strconv.ParseInt(v.Member.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, map[int64]int64{
			member: int64(v.Score),
		})
	}
	return data, nil
}
