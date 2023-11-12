package handler

import (
	"admin/model"
	"admin/storage"

	"github.com/mssola/user_agent"
)

type AdminLoginRecordHandler struct {
}

func (e *AdminLoginRecordHandler) AddRecord(adminId int64, clientIp string, userAgent string) {
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
	new(storage.AdminLoginRecordDb).AddRecord(record)
}
