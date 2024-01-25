package model

import (
	"core/model"
	"core/types"
)

type VerifyCodeRecord struct {
	Id         int64      `json:"id" xorm:"id pk BIGINT"`
	Username   string     `json:"username" xorm:"username"`
	SendType   int8       `json:"sendType" xorm:"send_type"`
	VerifyCode string     `json:"verifyCode" xorm:"verify_code"`
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
	ResCode    string     `json:"resCode" xorm:"res_code"`
	ResMsg     string     `json:"resMsg" xorm:"res_msg"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &VerifyCodeRecord{}

func (e *VerifyCodeRecord) TableName() string {
	return "verify_code_record"
}
