package listener

import (
	"business/storage"
	"core/event"
	"core/logger"
)

type NovelSubscribeListener struct {
	novelDb                *storage.NovelDb
	novelSubscribeNumCache *storage.NovelSubscribeNumCache
}

var _ event.Listener = &NovelSubscribeListener{}

func NewNovelSubscribeListener(novelDb *storage.NovelDb, novelSubscribeNumCache *storage.NovelSubscribeNumCache) *NovelSubscribeListener {
	return &NovelSubscribeListener{
		novelDb:                novelDb,
		novelSubscribeNumCache: novelSubscribeNumCache,
	}
}

func (e *NovelSubscribeListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	novelId := source.(int64)
	subscribeNum, err := e.novelSubscribeNumCache.Get(novelId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.novelDb.UpdateSubscribeNum(novelId, int32(subscribeNum))
}
