package handler

import (
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type RankHandler struct {
	ComicRankCache *storage.ComicRankCache
	NovelRankCache *storage.NovelRankCache
}

func (e *RankHandler) Rank(wc *web.WebContext) interface{} {
	rankTypeStr := wc.Query("rankType")
	if len(rankTypeStr) == 0 {
		return response.Success()
	}
	rankType, err := strconv.ParseInt(rankTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	timeTypeStr := wc.Query("timeType")
	if len(timeTypeStr) == 0 {
		wc.AbortWithError(err)
	}
	timeType, err := strconv.ParseInt(timeTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	assetTypeStr := wc.Query("assetType")
	if len(assetTypeStr) == 0 {
		return response.Success()
	}
	assetType, err := strconv.ParseInt(assetTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	wc.Info("rankType : %v, timeType : %v, assetType : %v, pageNum : %v, pageSize : %v", rankType, timeType, assetType, pageNum, pageSize)
	return response.Success()
}
