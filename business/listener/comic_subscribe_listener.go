package listener

import (
	"business/storage"
	"core/event"
	"core/logger"
)

type ComicSubscribeListener struct {
	comicDb                *storage.ComicDb
	comicSubscribeNumCache *storage.ComicSubscribeNumCache
}

var _ event.Listener = &ComicSubscribeListener{}

func NewComicSubscribeListener(comicDb *storage.ComicDb, comicSubscribeNumCache *storage.ComicSubscribeNumCache) *ComicSubscribeListener {
	return &ComicSubscribeListener{
		comicDb:                comicDb,
		comicSubscribeNumCache: comicSubscribeNumCache,
	}
}

func (e *ComicSubscribeListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	comicId := source.(int64)
	subscribeNum, err := e.comicSubscribeNumCache.Get(comicId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.comicDb.UpdateSubscribeNum(comicId, int32(subscribeNum))
}
