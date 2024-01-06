package listener

import (
	"business/storage"
	"business/transfer"
	"core/event"
	"core/logger"
)

type NovelHitListener struct {
	novelDb          *storage.NovelDb
	novelSubscribeDb *storage.NovelSubscribeDb
	novelHitNumCache *storage.NovelHitNumCache
}

var _ event.Listener = &NovelHitListener{}

func NewNovelHitListener(novelDb *storage.NovelDb, novelSubscribeDb *storage.NovelSubscribeDb, novelHitNumCache *storage.NovelHitNumCache) *NovelHitListener {
	return &NovelHitListener{
		novelDb:          novelDb,
		novelSubscribeDb: novelSubscribeDb,
		novelHitNumCache: novelHitNumCache,
	}
}

func (e *NovelHitListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	novelHitDto := source.(*transfer.NovelHitDto)
	novelId := novelHitDto.NovelId
	userId := novelHitDto.UserId
	hitNum, err := e.novelHitNumCache.Get(novelId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.novelDb.UpdateHitNum(novelId, int32(hitNum))
	if userId != 0 {
		e.novelSubscribeDb.SetRead(novelId, userId)
	}
}
