package handler

import (
	"basic/model"
	"basic/storage"
	"core/response"
	"core/web"
)

type CategoryGroupHandler struct {
	CategoryGroupDb *storage.CategoryGroupDb
}

func (e *CategoryGroupHandler) GetCategoryGroup(wc *web.WebContext) interface{} {
	categoryGroups, err := e.CategoryGroupDb.GetValidCategoryGroup()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(categoryGroups)
}

func (e *CategoryGroupHandler) GetGroupToCategory(wc *web.WebContext) interface{} {
	categoryGroups, err := e.CategoryGroupDb.GetValidCategoryGroup()
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(categoryGroups) == 0 {
		return response.Success()
	}
	var categories = make([]model.Category, 0)
	for _, v := range categoryGroups {
		categories = append(categories, model.Category{
			CategoryId:   v.CategoryGroupId,
			CategoryCode: v.GroupCode,
			CategoryName: v.GroupName,
			SeqNo:        v.SeqNo,
			Status:       v.Status,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return response.ReturnOK(categories)
}
