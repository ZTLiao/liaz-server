package handler

import (
	"business/enums"
	"business/resp"
	"business/storage"
	"core/redis"
	"core/response"
	"core/utils"
	"core/web"
	"sort"
	"strconv"
	"time"
)

type RankHandler struct {
	ComicDb            *storage.ComicDb
	NovelDb            *storage.NovelDb
	ComicRankCache     *storage.ComicRankCache
	NovelRankCache     *storage.NovelRankCache
	ComicRankItemCache *storage.ComicRankItemCache
	NovelRankItemCache *storage.NovelRankItemCache
}

func (e *RankHandler) Rank(wc *web.WebContext) interface{} {
	rankTypeStr := wc.Query("rankType")
	if len(rankTypeStr) == 0 {
		return response.Success()
	}
	rankType, err := strconv.ParseInt(rankTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	timeTypeStr := wc.Query("timeType")
	if len(timeTypeStr) == 0 {
		wc.AbortWithError(err)
	}
	timeType, err := strconv.ParseInt(timeTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	assetTypeStr := wc.Query("assetType")
	if len(assetTypeStr) == 0 {
		return response.Success()
	}
	assetType, err := strconv.ParseInt(assetTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var now = time.Now()
	var dateTime string
	if enums.TIME_TYPE_FOR_DAY == timeType {
		dateTime = now.Format(utils.NORM_DATE_PATTERN)
	} else if enums.TIME_TYPE_FOR_WEEK == timeType {
		dateTime = utils.GetStartOfWeek(now).Format(utils.NORM_DATE_PATTERN) + utils.COLON + utils.GetEndOfWeek(now).Format(utils.NORM_DATE_PATTERN)
	} else if enums.TIME_TYPE_FOR_MONTH == timeType {
		dateTime = now.Format(utils.NORM_MONTH_PATTERN)
	} else if enums.TIME_TYPE_FOR_TOTAL == timeType {
		dateTime = strconv.FormatInt(enums.TIME_TYPE_FOR_TOTAL, 10)
	}
	startIndex := (pageNum - 1) * pageSize
	stopIndex := startIndex + pageSize - 1
	wc.Info("rankType : %v, timeType : %v, dataTime : %v, assetType : %v, startIndex : %v, stopIndex : %v", rankType, timeType, dateTime, assetType, startIndex, stopIndex)
	var rankItems []resp.RankItemResp
	if enums.ASSET_TYPE_FOR_COMIC == assetType {
		rankItems, err = e.ComicRank(rankType, timeType, dateTime, int64(startIndex), int64(stopIndex))
	} else if enums.ASSET_TYPE_FOR_NOVEL == assetType {
		rankItems, err = e.NovelRank(rankType, timeType, dateTime, int64(startIndex), int64(stopIndex))
	}
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(rankItems)
}

func (e *RankHandler) ComicRank(rankType int64, timeType int64, dateTime string, startIndex int64, stopIndex int64) ([]resp.RankItemResp, error) {
	isExist, err := e.ComicRankItemCache.IsExist(rankType, timeType, dateTime)
	if err != nil {
		return nil, err
	}
	if !isExist {
		var redisLock = redis.NewRedisLock(e.ComicRankItemCache.RedisKey(rankType, timeType, dateTime))
		if !redisLock.Lock() {
			return nil, err
		}
		defer redisLock.Unlock()
		rankMap, err := e.ComicRankCache.Rank(rankType, timeType, dateTime, 0, 200)
		if err != nil {
			return nil, err
		}
		var comicIds = make([]int64, 0)
		for k := range rankMap {
			comicIds = append(comicIds, k)
		}
		if len(comicIds) == 0 {
			return nil, nil
		}
		comicMap, err := e.ComicDb.GetComicMapByIds(comicIds)
		if err != nil {
			return nil, err
		}
		var rankItems = make([]resp.RankItemResp, 0)
		for k, v := range rankMap {
			comic := comicMap[k]
			rankItems = append(rankItems, resp.RankItemResp{
				ObjId:      comic.ComicId,
				Title:      comic.Title,
				Cover:      comic.Cover,
				AssetType:  enums.ASSET_TYPE_FOR_COMIC,
				Categories: comic.Categories,
				Authors:    comic.Authors,
				Score:      v,
				UpdatedAt:  comic.EndTime,
			})
		}
		sort.Slice(rankItems, func(i, j int) bool {
			return rankItems[i].Score > rankItems[j].Score
		})
		for _, v := range rankItems {
			if v.ObjId != 0 {
				e.ComicRankItemCache.RPush(rankType, timeType, dateTime, v)
			}
		}
	}
	rankItems, err := e.ComicRankItemCache.LRange(rankType, timeType, dateTime, startIndex, stopIndex)
	if err != nil {
		return nil, err
	}
	return rankItems, nil
}

func (e *RankHandler) NovelRank(rankType int64, timeType int64, dateTime string, startIndex int64, stopIndex int64) ([]resp.RankItemResp, error) {
	isExist, err := e.NovelRankItemCache.IsExist(rankType, timeType, dateTime)
	if err != nil {
		return nil, err
	}
	if !isExist {
		var redisLock = redis.NewRedisLock(e.NovelRankItemCache.RedisKey(rankType, timeType, dateTime))
		if !redisLock.Lock() {
			return nil, err
		}
		defer redisLock.Unlock()
		rankMap, err := e.NovelRankCache.Rank(rankType, timeType, dateTime, 0, 200)
		if err != nil {
			return nil, err
		}
		var novelIds = make([]int64, 0)
		for k := range rankMap {
			novelIds = append(novelIds, k)
		}
		if len(novelIds) == 0 {
			return nil, nil
		}
		novelMap, err := e.NovelDb.GetNovelMapByIds(novelIds)
		if err != nil {
			return nil, err
		}
		var rankItems = make([]resp.RankItemResp, 0)
		for k, v := range rankMap {
			novel := novelMap[k]
			rankItems = append(rankItems, resp.RankItemResp{
				ObjId:      novel.NovelId,
				Title:      novel.Title,
				Cover:      novel.Cover,
				AssetType:  enums.ASSET_TYPE_FOR_NOVEL,
				Categories: novel.Categories,
				Authors:    novel.Authors,
				Score:      v,
				UpdatedAt:  novel.EndTime,
			})
		}
		sort.Slice(rankItems, func(i, j int) bool {
			return rankItems[i].Score > rankItems[j].Score
		})
		for _, v := range rankItems {
			e.NovelRankItemCache.RPush(rankType, timeType, dateTime, v)
		}
	}
	rankItems, err := e.NovelRankItemCache.LRange(rankType, timeType, dateTime, startIndex, stopIndex)
	if err != nil {
		return nil, err
	}
	return rankItems, nil
}
