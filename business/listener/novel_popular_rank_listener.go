package listener

import (
	"business/enums"
	"business/storage"
	"core/event"
	"core/utils"
	"strconv"
	"time"
)

type NovelPopularRankListener struct {
	novelRankCache *storage.NovelRankCache
}

var _ event.Listener = &NovelPopularRankListener{}

func NewNovelPopularRankListener(novelRankCache *storage.NovelRankCache) *NovelPopularRankListener {
	return &NovelPopularRankListener{novelRankCache}
}

func (e *NovelPopularRankListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	novelId := source.(int64)
	var now = time.Now()
	//日榜
	day := now.Format(utils.NORM_DATE_PATTERN)
	e.novelRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_DAY, day, novelId)
	//周榜
	week := utils.GetStartOfWeek(now).Format(utils.NORM_DATE_PATTERN) + utils.COLON + utils.GetEndOfWeek(now).Format(utils.NORM_DATE_PATTERN)
	e.novelRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, week, novelId)
	//月榜
	month := now.Format(utils.NORM_MONTH_PATTERN)
	e.novelRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_MONTH, month, novelId)
	//总榜
	e.novelRankCache.Incr(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_TOTAL, strconv.FormatInt(enums.TIME_TYPE_FOR_TOTAL, 10), novelId)
	//设置过期时间
	e.novelRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_DAY, day, time.Hour*24*2)
	e.novelRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, week, time.Hour*24*7*2)
	e.novelRankCache.Expire(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_MONTH, month, time.Hour*24*31*2)
}
