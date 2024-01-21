package model

import (
	"core/model"
	"core/types"
)

type MessageBoard struct {
	MessageBoardId int64      `json:"messageBoardId" xorm:"message_board_id pk BIGINT"`
	Content        string     `json:"content" xorm:"content"`
	DeviceId       string     `json:"deviceId" xorm:"device_id"`
	Os             string     `json:"os" xorm:"os"`
	OsVersion      string     `json:"osVersion" xorm:"os_version"`
	App            string     `json:"app" xorm:"app"`
	AppVersion     string     `json:"appVersion" xorm:"app_version"`
	Model          string     `json:"model" xorm:"model"`
	Imei           string     `json:"imei" xorm:"imei"`
	Channel        string     `json:"channel" xorm:"channel"`
	IspType        string     `json:"ispType" xorm:"isp_type"`
	NetType        string     `json:"netType" xorm:"net_type"`
	ClientIp       string     `json:"clientIp" xorm:"client_ip"`
	IpRegion       string     `json:"ipRegion" xorm:"ip_region"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestampz created"`
}

var _ model.BaseModel = &MessageBoard{}

func (e *MessageBoard) TableName() string {
	return "message_board"
}
