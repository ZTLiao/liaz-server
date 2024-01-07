package handler

import (
	"basic/device"
	"business/resp"
	"business/storage"
	"core/response"
	"core/utils"
	"core/web"
	"sort"
	"strconv"
	"strings"
)

type SearchHandler struct {
	SearchDb    *storage.SearchDb
	AssetDb     *storage.AssetDb
	SearchCache *storage.SearchCache
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
	searchs, err := e.AssetDb.Search(key, int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	var searchResps = make([]resp.SearchResp, 0)
	var result string
	if len(searchs) > 0 {
		var assetIds = make([]string, 0)
		for _, search := range searchs {
			assetId := search.AssetId
			searchResps = append(searchResps, resp.SearchResp{
				ObjId:          search.ObjId,
				Title:          search.Title,
				Cover:          search.Cover,
				AssetType:      search.AssetType,
				Categories:     search.Categories,
				Authors:        search.Authors,
				UpgradeChapter: search.UpgradeChapter,
			})
			assetIds = append(assetIds, strconv.FormatInt(assetId, 10))
			e.SearchCache.Incr(assetId)
		}
		result = strings.Join(assetIds, utils.COMMA)
	}
	deviceInfo := device.GetDeviceInfo(wc)
	userId := web.GetUserId(wc)
	e.SearchDb.SaveSearch(key, deviceInfo.DeviceId, userId, result)
	return response.ReturnOK(searchResps)
}

func (e *SearchHandler) HotRank(wc *web.WebContext) interface{} {
	members, err := e.SearchCache.Rank(0, 30)
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(members) == 0 {
		return response.Success()
	}
	var assetIds = make([]int64, 0)
	for k := range members {
		assetIds = append(assetIds, k)
	}
	assets, err := e.AssetDb.GetAssetByIds(assetIds)
	if err != nil {
		wc.AbortWithError(err)
	}
	sort.Slice(assets, func(i, j int) bool {
		return members[assets[i].AssetId] > members[assets[j].AssetId]
	})
	return response.ReturnOK(assets)
}
