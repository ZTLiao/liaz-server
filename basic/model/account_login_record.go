package model

import (
	"core/model"
	"core/types"
)

type AccountLoginRecord struct {
	Id         int64      `json:"id" xorm:"id pk autoincr BIGINT"`
	UserId     int64      `json:"userId" xorm:"user_id"`
	DeviceId   string     `json:"deviceId" xorm:"device_id"`
	Os         string     `json:"os" xorm:"os"`
	OsVersion  string     `json:"osVersion" xorm:"os_version"`
	App        string     `json:"app" xorm:"app"`
	AppVersion string     `json:"appVersion" xorm:"app_version"`
	Model      string     `json:"model" xorm:"model"`
	Imei       string     `json:"imei" xorm:"imei"`
	Channel    string     `json:"channel" xorm:"channel"`
	IspType    string     `json:"ispType" xorm:"isp_type"`
	NetType    string     `json:"netType" xorm:"net_type"`
	ClientIp   string     `json:"clientIp" xorm:"client_ip"`
	IpRegion   string     `json:"ipRegion" xorm:"ip_region"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &AccountLoginRecord{}

func (e *AccountLoginRecord) TableName() string {
	return "account_login_record"
}
