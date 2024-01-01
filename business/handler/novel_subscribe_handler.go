package handler

import (
	"business/storage"
	"core/constant"
	"core/response"
	"core/web"
	"strconv"
)

type NovelSubscribeHandler struct {
	NovelSubscribeDb *storage.NovelSubscribeDb
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
	} else {
		e.NovelSubscribeDb.DelNovelSubscribe(novelId, userId)
	}
	return response.Success()
}
