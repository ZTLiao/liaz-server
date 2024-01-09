package listener

import (
	"business/enums"
	"business/storage"
	"core/event"
	"core/utils"
	"strconv"
	"time"
)

type ComicPopularRankListener struct {
	comicRankCache *storage.ComicRankCache
}

var _ event.Listener = &ComicPopularRankListener{}

func NewComicPopularRankListener(comicRankCache *storage.ComicRankCache) *ComicPopularRankListener {
	return &ComicPopularRankListener{comicRankCache}
}

func (e *ComicPopularRankListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	comicId := source.(int64)
	var now = time.Now()
	//日榜
	day := now.Format(utils.NORM_DATE_PATTERN)
	e.comicRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_DAY, day, comicId)
	//周榜
	week := utils.GetStartOfWeek(now).Format(utils.NORM_DATE_PATTERN) + utils.COLON + utils.GetEndOfWeek(now).Format(utils.NORM_DATE_PATTERN)
	e.comicRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, week, comicId)
	//月榜
	month := now.Format(utils.NORM_MONTH_PATTERN)
	e.comicRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_MONTH, month, comicId)
	//总榜
	e.comicRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_TOTAL, strconv.FormatInt(enums.TIME_TYPE_FOR_TOTAL, 10), comicId)
	//设置过期时间
	e.comicRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_DAY, day, time.Hour*24*2)
	e.comicRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, week, time.Hour*24*7*2)
	e.comicRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_MONTH, month, time.Hour*24*31*2)
}
