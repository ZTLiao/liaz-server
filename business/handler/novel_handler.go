package handler

import (
	"business/storage"
	"core/response"
	"core/web"
)

type NovelHandler struct {
	NovelDb            *storage.NovelDb
	NovelChapterDb     *storage.NovelChapterDb
	NovelChapterItemDb *storage.NovelChapterItemDb
}

func (e *NovelHandler) NovelDetail(wc *web.WebContext) interface{} {
	return response.Success()
}

func (e *NovelHandler) NovelUpgrade(wc *web.WebContext) interface{} {
	return response.Success()
}
