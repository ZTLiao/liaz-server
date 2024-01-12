package storage

import (
	"business/resp"
	"core/constant"
	"core/redis"
	"encoding/json"
	"strconv"
	"time"
)

type NovelRankItemCache struct {
	redis *redis.RedisUtil
}

func NewNovelRankItemCache(redis *redis.RedisUtil) *NovelRankItemCache {
	return &NovelRankItemCache{redis}
}

func (e *NovelRankItemCache) RedisKey(rankType int64, timeType int64, dateTime string) string {
	return e.redis.GetKey(constant.NOVEL_RANK_ITEM, strconv.FormatInt(rankType, 10), strconv.FormatInt(timeType, 10), dateTime)
}

func (e *NovelRankItemCache) RPush(rankType int64, timeType int64, dateTime string, item resp.RankItemResp) error {
	var items = make([]string, 0)
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	items = append(items, string(data))
	_, err = e.redis.RPush(e.RedisKey(rankType, timeType, dateTime), items...)
	if err != nil {
		return err
	}
	err = e.redis.Expire(e.RedisKey(rankType, timeType, dateTime), time.Second*constant.TIME_OF_HOUR)
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelRankItemCache) LRange(rankType int64, timeType int64, dateTime string, start int64, stop int64) ([]resp.RankItemResp, error) {
	res, err := e.redis.LRange(e.RedisKey(rankType, timeType, dateTime), start, stop)
	if err != nil {
		return nil, err
	}
	var items []resp.RankItemResp
	for _, v := range res {
		var item resp.RankItemResp
		err = json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (e *NovelRankItemCache) IsExist(rankType int64, timeType int64, dateTime string) (bool, error) {
	num, err := e.redis.Exists(e.RedisKey(rankType, timeType, dateTime))
	if err != nil {
		return false, nil
	}
	return num > 0, nil
}
