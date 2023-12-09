package model

import "core/types"

type Device struct {
	DeviceId   string     `json:"deviceId" xorm:"device_id pk"`
	Os         string     `json:"os" xorm:"os"`
	OsVersion  string     `json:"osVersion" xorm:"os_version"`
	App        string     `json:"app" xorm:"app"`
	AppVersion string     `json:"appVersion" xorm:"app_version"`
	Model      string     `json:"model" xorm:"model"`
	Imei       string     `json:"imei" xorm:"imei"`
	Channel    string     `json:"channel" xorm:"channel"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt  types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *Device) TableName() string {
	return "device"
}
