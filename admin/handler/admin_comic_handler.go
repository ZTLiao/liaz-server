package handler

import (
	"admin/resp"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminComicHandler struct {
	ComicDb *storage.ComicDb
}

func (e *AdminComicHandler) GetComicPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	searchKey := wc.Query("searchKey")
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.ComicDb.GetComicPage(searchKey, pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}
