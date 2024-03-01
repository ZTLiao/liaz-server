package handler

import (
	"admin/resp"
	"basic/model"
	"basic/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminCategoryGroupHandler struct {
	CategoryGroupDb *storage.CategoryGroupDb
}

func (e *AdminCategoryGroupHandler) GetCategoryGroupPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.CategoryGroupDb.GetCategoryGroupPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminCategoryGroupHandler) GetCategoryGroupList(wc *web.WebContext) interface{} {
	categoryGroups, err := e.CategoryGroupDb.GetCategoryGroupList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(categoryGroups)
}

func (e *AdminCategoryGroupHandler) SaveCategoryGroup(wc *web.WebContext) interface{} {
	e.saveOrUpdateCategoryGroup(wc)
	return response.Success()
}

func (e *AdminCategoryGroupHandler) UpdateCategoryGroup(wc *web.WebContext) interface{} {
	e.saveOrUpdateCategoryGroup(wc)
	return response.Success()
}

func (e *AdminCategoryGroupHandler) saveOrUpdateCategoryGroup(wc *web.WebContext) {
	var categoryGroup = new(model.CategoryGroup)
	if err := wc.ShouldBindJSON(&categoryGroup); err != nil {
		wc.AbortWithError(err)
	}
	e.CategoryGroupDb.SaveOrUpdateCategoryGroup(categoryGroup)
}

func (e *AdminCategoryGroupHandler) DelCategoryGroup(wc *web.WebContext) interface{} {
	categoryIdStr := wc.Param("categoryGroupId")
	if len(categoryIdStr) > 0 {
		categoryGroupId, err := strconv.ParseInt(categoryIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.CategoryGroupDb.DelCategoryGroup(categoryGroupId)
	}
	return response.Success()
}
