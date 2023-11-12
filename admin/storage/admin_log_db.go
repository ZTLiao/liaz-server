package storage

import (
	"admin/model"
	"core/application"
	"encoding/json"
)

type AdminLogDb struct {
}

func (e *AdminLogDb) AddLog(adminId int64, uri string, headers map[string][]string, queryParams map[string]any, formParams map[string]any, bodyParams string) {
	var adminLog = new(model.AdminLog)
	adminLog.AdminId = adminId
	adminLog.Uri = uri
	headersJson, _ := json.Marshal(headers)
	queryParamsJson, _ := json.Marshal(queryParams)
	formParamsJson, _ := json.Marshal(formParams)
	var params = &map[string]string{
		"headers":     string(headersJson),
		"queryParams": string(queryParamsJson),
		"formParams":  string(formParamsJson),
		"bodyParams":  bodyParams,
	}
	paramsJson, _ := json.Marshal(params)
	adminLog.Params = string(paramsJson)
	var engine = application.GetApp().GetXormEngine()
	engine.Insert(adminLog)
}
