package handler

import (
	basicStorage "basic/storage"
	"business/resp"
	businessStorage "business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type CategorySearchHandler struct {
	CategoryGroupDb *basicStorage.CategoryGroupDb
	CategoryDb      *basicStorage.CategoryDb
	AssetDb         *businessStorage.AssetDb
}

func (e *CategorySearchHandler) GetCategorySearch(wc *web.WebContext) interface{} {
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
	categoryGroup, err := e.CategoryGroupDb.GetCategoryGroupById(categoryId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if categoryGroup == nil {
		return response.Success()
	}
	groupCode := categoryGroup.GroupCode
	if len(groupCode) == 0 {
		return response.Success()
	}
	categories, err := e.CategoryDb.GetCategoryByGroupCode(groupCode)
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(categories) == 0 {
		return response.Success()
	}
	var categoryIds = make([]int64, 0)
	for _, v := range categories {
		categoryIds = append(categoryIds, v.CategoryId)
	}
	assets, err := e.AssetDb.GetAssetByCategoryIds(int8(assetType), categoryIds, int32(pageNum), int32(pageSize))
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
			ObjId:          asset.ObjId,
			ChapterId:      asset.ChapterId,
			UpgradeChapter: asset.UpgradeChapter,
		})
	}
	return response.ReturnOK(categoryItems)
}
