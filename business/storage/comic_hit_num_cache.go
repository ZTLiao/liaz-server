package storage

import (
	"core/constant"
	"core/redis"
	"strconv"
)

type ComicHitNumCache struct {
	redis *redis.RedisUtil
}

func NewComicHitNumCache(redis *redis.RedisUtil) *ComicHitNumCache {
	return &ComicHitNumCache{redis}
}

func (e *ComicHitNumCache) RedisKey() string {
	return e.redis.GetKey(constant.COMIC_HIT_NUM)
}

func (e *ComicHitNumCache) Incr(comicId int64) (int64, error) {
	score, err := e.redis.ZIncrBy(e.RedisKey(), 1, strconv.FormatInt(comicId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *ComicHitNumCache) Decr(comicId int64) error {
	_, err := e.redis.ZIncrBy(e.RedisKey(), -1, strconv.FormatInt(comicId, 10))
	if err != nil {
		return err
	}
	return nil
}

func (e *ComicHitNumCache) Get(comicId int64) (int64, error) {
	score, err := e.redis.ZScore(e.RedisKey(), strconv.FormatInt(comicId, 10))
	if err != nil {
		return 0, err
	}
	return int64(score), nil
}

func (e *ComicHitNumCache) Rank(startIndex int64, stopIndex int64) ([]map[int64]int64, error) {
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
