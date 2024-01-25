package model

import (
	"core/model"
	"core/types"
)

type Search struct {
	SearchId  int64      `json:"searchId" xorm:"search_id pk autoincr BIGINT"`
	SearchKey string     `json:"searchKey" xorm:"search_key"`
	DeviceId  string     `json:"deviceId" xorm:"device_id"`
	UserId    int64      `json:"userId" xorm:"user_id"`
	Result    string     `json:"result" xorm:"result"`
	Status    int8       `json:"status" xorm:"status"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &Search{}

func (e *Search) TableName() string {
	return "search"
}
