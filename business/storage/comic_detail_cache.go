package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"strconv"
)

type ComicDetailCache struct {
	redis *redis.RedisUtil
}

func NewComicDetailCache(redis *redis.RedisUtil) *ComicDetailCache {
	return &ComicDetailCache{redis}
}

func (e *ComicDetailCache) RedisKey(comicId int64) string {
	return e.redis.GetKey(constant.COMIC_DETAIL, strconv.FormatInt(comicId, 10))
}

func (e *ComicDetailCache) Set(comicId int64, comicDetail *resp.ComicDetailResp) error {
	if comicDetail == nil {
		return nil
	}
	data, err := json.Marshal(comicDetail)
	if err != nil {
		return err
	}
	return e.redis.Set(e.RedisKey(comicId), data, constant.TIME_OF_HALF_DAY)
}

func (e *ComicDetailCache) Get(comicId int64) (*resp.ComicDetailResp, error) {
	data, err := e.redis.Get(e.RedisKey(comicId))
	if err != nil {
		return nil, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	var comicDetail resp.ComicDetailResp
	err = json.Unmarshal([]byte(data), &comicDetail)
	if err != nil {
		return nil, err
	}
	return &comicDetail, nil
}

func (e *ComicDetailCache) Del(comicId int64) error {
	return e.redis.Del(e.RedisKey(comicId))
}

func (e *ComicDetailCache) IsExist(comicId int64) (bool, error) {
	num, err := e.redis.Exists(e.RedisKey(comicId))
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}
