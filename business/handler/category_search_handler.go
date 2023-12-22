package handler

import (
	"business/enums"
	"core/response"
	"core/web"
	"strconv"
)

type CategorySearchHandler struct {
	ComicHandler *ComicHandler
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
	if enums.ASSET_TYPE_FOR_ALL == assetType {

	} else if enums.ASSET_TYPE_FOR_COMIC == assetType {
		return e.ComicHandler.GetComicByCategory(wc)
	} else if enums.ASSET_TYPE_FOR_NOVEL == assetType {

	}
	return response.Success()
}
