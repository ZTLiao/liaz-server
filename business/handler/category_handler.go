package handler

import (
	"basic/storage"
	"core/response"
	"core/web"
)

type CategoryHandler struct {
	CategoryDb *storage.CategoryDb
}

func (e *CategoryHandler) GetCategory(wc *web.WebContext) interface{} {
	categories, err := e.CategoryDb.GetValidCategory()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(categories)
}
