package handler

import (
	"admin/resp"
	"business/model"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminRecommendItemHandler struct {
	RecommendItemDb *storage.RecommendItemDb
}

func (e *AdminRecommendItemHandler) GetRecommendItemPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	recommendId, err := strconv.ParseInt(wc.DefaultQuery("recommendId", "0"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.RecommendItemDb.GetRecommendItemPage(recommendId, pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminRecommendItemHandler) SaveRecommendItem(wc *web.WebContext) interface{} {
	e.saveOrUpdateRecommendItem(wc)
	return response.Success()
}

func (e *AdminRecommendItemHandler) UpdateRecommendItem(wc *web.WebContext) interface{} {
	e.saveOrUpdateRecommendItem(wc)
	return response.Success()
}

func (e *AdminRecommendItemHandler) saveOrUpdateRecommendItem(wc *web.WebContext) {
	var recommendItem = new(model.RecommendItem)
	if err := wc.ShouldBindJSON(&recommendItem); err != nil {
		wc.AbortWithError(err)
	}
	e.RecommendItemDb.SaveOrUpdateRecommendItem(recommendItem)
}

func (e *AdminRecommendItemHandler) DelRecommendItem(wc *web.WebContext) interface{} {
	recommendItemIdStr := wc.Param("recommendItemId")
	if len(recommendItemIdStr) > 0 {
		recommendItemId, err := strconv.ParseInt(recommendItemIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.RecommendItemDb.DelRecommendItem(recommendItemId)
	}
	return response.Success()
}
