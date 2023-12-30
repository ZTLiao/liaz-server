package handler

import (
	"basic/handler"
	"basic/model"
	"core/response"
	"core/web"
	"fmt"
	"strconv"
)

type AdminSysConfHandler struct {
	SysConfHandler *handler.SysConfHandler
}

func (e *AdminSysConfHandler) GetAdminSysConf(wc *web.WebContext) interface{} {
	sysConfs, err := e.SysConfHandler.GetSysConfList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(sysConfs)
}

func (e *AdminSysConfHandler) SaveAdminSysConf(wc *web.WebContext) interface{} {
	e.SaveOrUpdateAdminSysConf(wc)
	return response.Success()
}

func (e *AdminSysConfHandler) UpdateAdminSysConf(wc *web.WebContext) interface{} {
	e.SaveOrUpdateAdminSysConf(wc)
	return response.Success()
}

func (e *AdminSysConfHandler) DelAdminSysConf(wc *web.WebContext) interface{} {
	confIdStr := wc.Param("confId")
	if len(confIdStr) > 0 {
		confId, err := strconv.ParseInt(confIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.SysConfHandler.DeleteSysConf(confId)
	}
	return response.Success()
}

func (e *AdminSysConfHandler) SaveOrUpdateAdminSysConf(wc *web.WebContext) {
	var params map[string]any
	if err := wc.ShouldBindJSON(&params); err != nil {
		wc.AbortWithError(err)
	}
	confIdStr := fmt.Sprint(params["confId"])
	confKey := fmt.Sprint(params["confKey"])
	confName := fmt.Sprint(params["confName"])
	confKindStr := fmt.Sprint(params["confKind"])
	confTypeStr := fmt.Sprint(params["confType"])
	confValue := fmt.Sprint(params["confValue"])
	statusStr := fmt.Sprint(params["status"])
	description := fmt.Sprint(params["description"])
	var sysConf = new(model.SysConf)
	if len(confIdStr) > 0 {
		confId, err := strconv.ParseInt(confIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		sysConf.ConfId = confId
	}
	sysConf.ConfKey = confKey
	sysConf.ConfName = confName
	confKind, err := strconv.ParseInt(confKindStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	sysConf.ConfKind = int8(confKind)
	confType, err := strconv.ParseInt(confTypeStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	sysConf.ConfType = int8(confType)
	sysConf.ConfValue = confValue
	status, err := strconv.ParseInt(statusStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	sysConf.Status = int8(status)
	sysConf.Description = description
	if sysConf.ConfId == 0 {
		e.SysConfHandler.SaveSysConf(sysConf)
	} else {
		e.SysConfHandler.UpdateSysConf(sysConf)
	}
}

func (e *AdminSysConfHandler) GetAdminSysConfByType(wc *web.WebContext) interface{} {
	confTypeStr := wc.Param("confType")
	if len(confTypeStr) == 0 {
		return response.Success()
	}
	confType, err := strconv.ParseInt(confTypeStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	sysConfs, err := e.SysConfHandler.GetSysConfByType(int8(confType))
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(sysConfs)
}
