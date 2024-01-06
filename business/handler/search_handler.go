package handler

import (
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type SearchHandler struct {
	AssetDb *storage.AssetDb
}

func (e *SearchHandler) Search(wc *web.WebContext) interface{} {
	key := wc.Query("key")
	if len(key) == 0 {
		return response.Success()
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	_, err = e.AssetDb.Search(key, int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.Success()
}
