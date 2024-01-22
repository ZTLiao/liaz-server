package handler

import (
	"admin/resp"
	"business/model"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type AdminAppVersionHandler struct {
	AppVersionDb *storage.AppVersionDb
}

func (e *AdminAppVersionHandler) GetAppVersionPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.AppVersionDb.GetAppVersionPage(pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminAppVersionHandler) SaveAppVersion(wc *web.WebContext) interface{} {
	e.saveOrUpdateAppVersion(wc)
	return response.Success()
}

func (e *AdminAppVersionHandler) UpdateAppVersion(wc *web.WebContext) interface{} {
	e.saveOrUpdateAppVersion(wc)
	return response.Success()
}

func (e *AdminAppVersionHandler) saveOrUpdateAppVersion(wc *web.WebContext) {
	var appVersion = new(model.AppVersion)
	if err := wc.ShouldBindJSON(&appVersion); err != nil {
		wc.AbortWithError(err)
	}
	e.AppVersionDb.SaveOrUpdateAppVersion(appVersion)
}

func (e *AdminAppVersionHandler) DelAppVersion(wc *web.WebContext) interface{} {
	versionIdStr := wc.Param("versionId")
	if len(versionIdStr) > 0 {
		versionId, err := strconv.ParseInt(versionIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AppVersionDb.DelAppVersion(versionId)
	}
	return response.Success()
}
