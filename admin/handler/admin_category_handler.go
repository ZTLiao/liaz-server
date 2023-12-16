package handler

import (
	"admin/resp"
	"basic/model"
	"basic/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminCategoryHandler struct {
	CategoryDb *storage.CategoryDb
}

func (e *AdminCategoryHandler) GetCategoryPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.CategoryDb.GetCategoryPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminCategoryHandler) GetCategoryList(wc *web.WebContext) interface{} {
	categorys, err := e.CategoryDb.GetCategoryList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(categorys)
}

func (e *AdminCategoryHandler) SaveCategory(wc *web.WebContext) interface{} {
	e.saveOrUpdateCategory(wc)
	return response.Success()
}

func (e *AdminCategoryHandler) UpdateCategory(wc *web.WebContext) interface{} {
	e.saveOrUpdateCategory(wc)
	return response.Success()
}

func (e *AdminCategoryHandler) saveOrUpdateCategory(wc *web.WebContext) {
	var category = new(model.Category)
	if err := wc.ShouldBindJSON(&category); err != nil {
		wc.AbortWithError(err)
	}
	e.CategoryDb.SaveOrUpdateCategory(category)
}

func (e *AdminCategoryHandler) DelCategory(wc *web.WebContext) interface{} {
	categoryIdStr := wc.Param("categoryId")
	if len(categoryIdStr) > 0 {
		categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.CategoryDb.DelCategory(categoryId)
	}
	return response.Success()
}
