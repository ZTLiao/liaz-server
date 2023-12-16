package handler

import (
	"admin/resp"
	"basic/model"
	"basic/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminAuthorHandler struct {
	AuthorDb *storage.AuthorDb
}

func (e *AdminAuthorHandler) GetAuthorPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.AuthorDb.GetAuthorPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminAuthorHandler) GetAuthorList(wc *web.WebContext) interface{} {
	categorys, err := e.AuthorDb.GetAuthorList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(categorys)
}

func (e *AdminAuthorHandler) SaveAuthor(wc *web.WebContext) interface{} {
	e.saveOrUpdateAuthor(wc)
	return response.Success()
}

func (e *AdminAuthorHandler) UpdateAuthor(wc *web.WebContext) interface{} {
	e.saveOrUpdateAuthor(wc)
	return response.Success()
}

func (e *AdminAuthorHandler) saveOrUpdateAuthor(wc *web.WebContext) {
	var author = new(model.Author)
	if err := wc.ShouldBindJSON(&author); err != nil {
		wc.AbortWithError(err)
	}
	e.AuthorDb.SaveOrUpdateAuthor(author)
}

func (e *AdminAuthorHandler) DelAuthor(wc *web.WebContext) interface{} {
	authorIdStr := wc.Param("authorId")
	if len(authorIdStr) > 0 {
		authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AuthorDb.DelAuthor(authorId)
	}
	return response.Success()
}
