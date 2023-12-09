package model

import "core/types"

type ClientInitRecord struct {
	Id         int64      `json:"id" xorm:"id pk BIGINT"`
	DeviceId   string     `json:"deviceId" xorm:"device_id pk"`
	Os         string     `json:"os" xorm:"os"`
	OsVersion  string     `json:"osVersion" xorm:"os_version"`
	App        string     `json:"app" xorm:"app"`
	AppVersion string     `json:"appVersion" xorm:"app_version"`
	Model      string     `json:"model" xorm:"model"`
	Imei       string     `json:"imei" xorm:"imei"`
	Channel    string     `json:"channel" xorm:"channel"`
	IspType    string     `json:"ispType" xorm:"isp_type"`
	NetType    string     `json:"netType" xorm:"net_type"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at"`
}

func (e *ClientInitRecord) TableName() string {
	return "client_init_record"
}
