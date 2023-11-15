package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
	"core/types"

	"github.com/mssola/user_agent"
)

type AdminLoginRecordDb struct {
}

func (e *AdminLoginRecordDb) AddRecord(adminId int64, clientIp string, userAgent string) {
	ua := user_agent.New(userAgent)
	var record = new(model.AdminLoginRecord)
	record.AdminId = adminId
	record.IsMobile = ua.Mobile()
	record.IsBot = ua.Bot()
	record.Mozilla = ua.Mozilla()
	record.Platform = ua.Platform()
	record.Os = ua.OS()
	engineName, engineVersion := ua.Engine()
	record.EngineName = engineName
	record.EngineVersion = engineVersion
	browserName, browserVersion := ua.Browser()
	record.BrowserName = browserName
	record.BrowserVersion = browserVersion
	record.ClientIp = clientIp
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
