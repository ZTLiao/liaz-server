package handler

import (
	"business/storage"
	"core/response"
	"core/web"
)

type DiscussThumbHandler struct {
	DiscussThumbDb *storage.DiscussThumbDb
}

func (e *DiscussThumbHandler) Thumb(wc *web.WebContext) interface{} {
	return response.Success()
}
