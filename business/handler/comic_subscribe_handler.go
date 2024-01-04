package handler

import (
	"business/storage"
	"core/constant"
	"core/event"
	"core/response"
	"core/web"
	"strconv"
)

type ComicSubscribeHandler struct {
	ComicSubscribeDb       *storage.ComicSubscribeDb
	ComicSubscribeNumCache *storage.ComicSubscribeNumCache
}

func (e *ComicSubscribeHandler) Subscribe(wc *web.WebContext) interface{} {
	comicIdStr := wc.PostForm("comicId")
	isSubscribeStr := wc.PostForm("isSubscribe")
	comicId, err := strconv.ParseInt(comicIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	isSubscribe, err := strconv.ParseInt(isSubscribeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	if int8(isSubscribe) == constant.YES {
		e.ComicSubscribeDb.SaveComicSubscribe(comicId, userId)
		e.ComicSubscribeNumCache.Incr(comicId)
	} else {
		e.ComicSubscribeDb.DelComicSubscribe(comicId, userId)
		e.ComicSubscribeNumCache.Decr(comicId)
	}
	event.Bus.Publish(constant.COMIC_SUBSCRIBE_TOPIC, comicId)
	return response.Success()
}
