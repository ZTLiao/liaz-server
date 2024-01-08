package handler

import (
	"business/storage"
	"core/constant"
	"core/event"
	"core/response"
	"core/web"
	"strconv"
)

type NovelSubscribeHandler struct {
	NovelSubscribeDb       *storage.NovelSubscribeDb
	NovelSubscribeNumCache *storage.NovelSubscribeNumCache
}

func (e *NovelSubscribeHandler) Subscribe(wc *web.WebContext) interface{} {
	novelIdStr := wc.PostForm("novelId")
	isSubscribeStr := wc.PostForm("isSubscribe")
	novelId, err := strconv.ParseInt(novelIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	isSubscribe, err := strconv.ParseInt(isSubscribeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	if int8(isSubscribe) == constant.YES {
		e.NovelSubscribeDb.SaveNovelSubscribe(novelId, userId)
		e.NovelSubscribeNumCache.Incr(novelId)
		event.Bus.Publish(constant.NOVEL_SUBSCRIBE_RANK_TOPIC, novelId)
	} else {
		e.NovelSubscribeDb.DelNovelSubscribe(novelId, userId)
		e.NovelSubscribeNumCache.Decr(novelId)
	}
	event.Bus.Publish(constant.NOVEL_SUBSCRIBE_TOPIC, novelId)
	return response.Success()
}
