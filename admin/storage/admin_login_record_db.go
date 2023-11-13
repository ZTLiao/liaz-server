package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
	"core/types"
)

type AdminLoginRecordDb struct {
}

func (e *AdminLoginRecordDb) AddRecord(record *model.AdminLoginRecord) {
	var engine = application.GetXormEngine()
	if _, err := engine.Insert(record); err != nil {
		logger.Error(err.Error())
	}
}

func (e *AdminLoginRecordDb) GetLastTime(adminId int64) types.Time {
	var engine = application.GetXormEngine()
	var record = new(model.AdminLoginRecord)
	has, err := engine.Where("admin_id = ?", adminId).OrderBy("record_id desc").Limit(1, 1).Get(record)
	if err != nil {
		logger.Error(err.Error())
	}
	if !has {
		return types.Time{}
	}
	return record.CreatedAt
}
