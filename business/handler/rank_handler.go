package handler

import (
	"business/enums"
	"business/resp"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type RankHandler struct {
	ComicDb        *storage.ComicDb
	NovelDb        *storage.NovelDb
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
	var rankItems []resp.RankItemResp
	if enums.ASSET_TYPE_FOR_COMIC == assetType {
		rankItems, err = e.ComicRank(rankType, timeType, int32(pageNum), int32(pageSize))
	} else if enums.ASSET_TYPE_FOR_NOVEL == assetType {
		rankItems, err = e.NovelRank(rankType, timeType, int32(pageNum), int32(pageSize))
	}
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(rankItems)
}

func (e *RankHandler) ComicRank(rankType int64, timeType int64, pageNum int32, pageSize int32) ([]resp.RankItemResp, error) {
	return nil, nil
}

func (e *RankHandler) NovelRank(rankType int64, timeType int64, pageNum int32, pageSize int32) ([]resp.RankItemResp, error) {
	return nil, nil
}
