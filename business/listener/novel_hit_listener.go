package listener

import (
	"business/storage"
	"core/event"
	"core/logger"
)

type NovelHitListener struct {
	novelDb          *storage.NovelDb
	novelHitNumCache *storage.NovelHitNumCache
}

var _ event.Listener = &NovelHitListener{}

func NewNovelHitListener(novelDb *storage.NovelDb, novelHitNumCache *storage.NovelHitNumCache) *NovelHitListener {
	return &NovelHitListener{
		novelDb:          novelDb,
		novelHitNumCache: novelHitNumCache,
	}
}

func (e *NovelHitListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	novelId := source.(int64)
	hitNum, err := e.novelHitNumCache.Get(novelId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.novelDb.UpdateHitNum(novelId, int32(hitNum))
}
