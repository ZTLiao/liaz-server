package listener

import (
	"business/storage"
	"business/transfer"
	"core/event"
	"core/logger"
)

type ComicHitListener struct {
	comicDb          *storage.ComicDb
	comicSubscribeDb *storage.ComicSubscribeDb
	comicHitNumCache *storage.ComicHitNumCache
}

var _ event.Listener = &ComicHitListener{}

func NewComicHitListener(comicDb *storage.ComicDb, comicSubscribeDb *storage.ComicSubscribeDb, comicHitNumCache *storage.ComicHitNumCache) *ComicHitListener {
	return &ComicHitListener{
		comicDb:          comicDb,
		comicSubscribeDb: comicSubscribeDb,
		comicHitNumCache: comicHitNumCache,
	}
}

func (e *ComicHitListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	comicHitDto := source.(*transfer.ComicHitDto)
	comicId := comicHitDto.ComicId
	userId := comicHitDto.UserId
	hitNum, err := e.comicHitNumCache.Get(comicId)
	if err != nil {
		logger.Panic(err.Error())
	}
	e.comicDb.UpdateHitNum(comicId, int32(hitNum))
	if userId != 0 {
		e.comicSubscribeDb.SetRead(comicId, userId)
	}
}
