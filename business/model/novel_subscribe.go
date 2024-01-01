package model

import (
	"core/model"
	"core/types"
)

type NovelSubscribe struct {
	NovelSubscribeId int64      `json:"novelSubscribeId" xorm:"novel_subscribe_id pk autoincr BIGINT"`
	UserId           int64      `json:"userId" xorm:"user_id"`
	NovelId          int64      `json:"novelId" xorm:"novel_id"`
	CreatedAt        types.Time `json:"createdAt" xorm:"created_at timestampz created"`
}

var _ model.BaseModel = &NovelSubscribe{}

func (e *NovelSubscribe) TableName() string {
	return "novel_subscribe"
}
