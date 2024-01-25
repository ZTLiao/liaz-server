package model

import (
	"core/model"
	"core/types"
)

type AdminLoginRecord struct {
	RecordId       int64      `json:"recordId" xorm:"record_id pk autoincr BIGINT"`
	AdminId        int64      `json:"AdminId" xorm:"admin_id"`
	IsMobile       bool       `json:"isMobile" xorm:"is_mobile"`
	IsBot          bool       `json:"isBot" xorm:"is_bot"`
	Mozilla        string     `json:"mozilla" xorm:"mozilla"`
	Platform       string     `json:"platform" xorm:"platform"`
	Os             string     `json:"os" xorm:"os"`
	EngineName     string     `json:"engineName" xorm:"engine_name"`
	EngineVersion  string     `json:"engineVersion" xorm:"engine_version"`
	BrowserName    string     `json:"browserName" xorm:"broswer_name"`
	BrowserVersion string     `json:"browserVersion" xorm:"broswer_version"`
	ClientIp       string     `json:"clientIp" xorm:"client_ip"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &AdminLoginRecord{}

func (e *AdminLoginRecord) TableName() string {
	return "admin_login_record"
}
