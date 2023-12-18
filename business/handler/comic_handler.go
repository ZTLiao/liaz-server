package handler

import (
	"business/storage"
	"core/response"
	"core/web"
)

type ComicHandler struct {
	ComicDb            *storage.ComicDb
	ComicChapterDb     *storage.ComicChapterDb
	ComicChapterItemDb *storage.ComicChapterItemDb
}

func (e *ComicHandler) GetComicDetail(wc *web.WebContext) interface{} {
	return response.ReturnOK(nil)
}
