package listener

import (
	"business/storage"
	"core/event"
	"core/logger"
)

type ComicHitListener struct {
	comicDb          *storage.ComicDb
	comicHitNumCache *storage.ComicHitNumCache
}

var _ event.Listener = &ComicHitListener{}

func NewComicHitListener(comicDb *storage.ComicDb, comicHitNumCache *storage.ComicHitNumCache) *ComicHitListener {
	return &ComicHitListener{
		comicDb:          comicDb,
		comicHitNumCache: comicHitNumCache,
	}
}

func (e *ComicHitListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	comicId := source.(int64)
	hitNum, err := e.comicHitNumCache.Get(comicId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.comicDb.UpdateHitNum(comicId, int32(hitNum))
}
