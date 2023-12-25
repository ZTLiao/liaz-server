package storage

import (
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/mssola/user_agent"
)

type AdminLoginRecordDb struct {
	db *xorm.Engine
}

func NewAdminLoginRecordDb(db *xorm.Engine) *AdminLoginRecordDb {
	return &AdminLoginRecordDb{db}
}

func (e *AdminLoginRecordDb) AddRecord(adminId int64, clientIp string, userAgent string) error {
	var now = types.Time(time.Now())
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
	record.CreatedAt = now
	if _, err := e.db.Insert(record); err != nil {
		return err
	}
	return nil
}

func (e *AdminLoginRecordDb) GetLastTime(adminId int64) (types.Time, error) {
	var record = new(model.AdminLoginRecord)
	_, err := e.db.Where("admin_id = ?", adminId).OrderBy("record_id desc").Limit(1, 1).Get(record)
	if err != nil {
		return types.Time{}, err
	}
	return record.CreatedAt, nil
}
