package model

import "core/types"

type UserDevice struct {
	UserId    int64      `json:"userId" xorm:"user_id"`
	DeviceId  string     `json:"deviceId" xorm:"device_id"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *UserDevice) TableName() string {
	return "user_device"
}
