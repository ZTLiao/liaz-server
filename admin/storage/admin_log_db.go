package storage

import (
	"admin/model"
	"encoding/json"

	"github.com/go-xorm/xorm"
)

type AdminLogDb struct {
	db *xorm.Engine
}

func NewAdminLogDb(db *xorm.Engine) *AdminLogDb {
	return &AdminLogDb{db}
}

func (e *AdminLogDb) AddLog(adminId int64, uri string, headers map[string][]string, queryParams map[string]any, formParams map[string]any, bodyParams string) error {
	var adminLog = new(model.AdminLog)
	adminLog.AdminId = adminId
	adminLog.Uri = uri
	headersJson, err := json.Marshal(headers)
	if err != nil {
		return err
	}
	queryParamsJson, err := json.Marshal(queryParams)
	if err != nil {
		return err
	}
	formParamsJson, err := json.Marshal(formParams)
	if err != nil {
		return err
	}
	var params = &map[string]string{
		"headers":     string(headersJson),
		"queryParams": string(queryParamsJson),
		"formParams":  string(formParamsJson),
		"bodyParams":  bodyParams,
	}
	paramsJson, err := json.Marshal(params)
	if err != nil {
		return err
	}
	adminLog.Params = string(paramsJson)
	_, err = e.db.Insert(adminLog)
	if err != nil {
		return err
	}
	return nil
}
