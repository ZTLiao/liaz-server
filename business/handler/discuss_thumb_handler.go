package handler

import (
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type DiscussThumbHandler struct {
	DiscussThumbDb       *storage.DiscussThumbDb
	DiscussThumbNumCache *storage.DiscussThumbNumCache
}

func (e *DiscussThumbHandler) Thumb(wc *web.WebContext) interface{} {
	discussIdStr := wc.PostForm("discussId")
	if len(discussIdStr) == 0 {
		return response.Success()
	}
	discussId, err := strconv.ParseInt(discussIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	err = e.DiscussThumbDb.Save(discussId, userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	e.DiscussThumbNumCache.Incr(discussId)
	return response.Success()
}
