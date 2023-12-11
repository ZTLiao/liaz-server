package handler

import (
	"business/storage"
	"core/response"
	"core/web"
	"fmt"
	"strconv"
)

type RecommendHandler struct {
	RecommendDb     *storage.RecommendDb
	RecommendItemDb *storage.RecommendItemDb
	RecommendCache  *storage.RecommendCache
}

func (e *RecommendHandler) Recommend(wc *web.WebContext) interface{} {
	positionStr := wc.Param("position")
	var position int64
	var err error
	if len(positionStr) > 0 {
		position, err = strconv.ParseInt(positionStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	fmt.Println(position)
	return response.Success()
}
