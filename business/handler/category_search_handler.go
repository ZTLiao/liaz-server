package handler

import (
	"business/resp"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type CategorySearchHandler struct {
	AssetDb *storage.AssetDb
}

func (e *CategorySearchHandler) GetContent(wc *web.WebContext) interface{} {
	assetTypeStr := wc.Query("assetType")
	if len(assetTypeStr) == 0 {
		return response.Success()
	}
	assetType, err := strconv.ParseInt(assetTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	categoryIdStr := wc.Query("categoryId")
	if len(categoryIdStr) == 0 {
		return response.Success()
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
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
	assets, err := e.AssetDb.GetAssetByCategoryId(int8(assetType), categoryId, int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(assets) == 0 {
		return response.Success()
	}
	var categoryItems = make([]resp.CategoryItemResp, 0)
	for _, asset := range assets {
		categoryItems = append(categoryItems, resp.CategoryItemResp{
			CategoryId:     categoryId,
			AssetType:      asset.AssetType,
			Title:          asset.Title,
			Cover:          asset.Cover,
			UpgradeChapter: asset.UpgradeChapter,
			ObjId:          asset.ObjId,
		})
	}
	return response.ReturnOK(categoryItems)
}
