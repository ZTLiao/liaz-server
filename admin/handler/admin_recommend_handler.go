package handler

import (
	"admin/resp"
	"business/model"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminRecommendHandler struct {
	RecommendDb *storage.RecommendDb
}

func (e *AdminRecommendHandler) GetRecommendPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.RecommendDb.GetRecommendPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminRecommendHandler) GetRecommendList(wc *web.WebContext) interface{} {
	recommends, err := e.RecommendDb.GetRecommendList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(recommends)
}

func (e *AdminRecommendHandler) SaveRecommend(wc *web.WebContext) interface{} {
	e.saveOrUpdateRecommend(wc)
	return response.Success()
}

func (e *AdminRecommendHandler) UpdateRecommend(wc *web.WebContext) interface{} {
	e.saveOrUpdateRecommend(wc)
	return response.Success()
}

func (e *AdminRecommendHandler) saveOrUpdateRecommend(wc *web.WebContext) {
	var recommend = new(model.Recommend)
	if err := wc.ShouldBindJSON(&recommend); err != nil {
		wc.AbortWithError(err)
	}
	e.RecommendDb.SaveOrUpdateRecommend(recommend)
}

func (e *AdminRecommendHandler) DelRecommend(wc *web.WebContext) interface{} {
	recommendIdStr := wc.Param("recommendId")
	if len(recommendIdStr) > 0 {
		recommendId, err := strconv.ParseInt(recommendIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.RecommendDb.DelRecommend(recommendId)
	}
	return response.Success()
}
