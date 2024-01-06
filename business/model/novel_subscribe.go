package model

import (
	"core/model"
	"core/types"
)

type NovelSubscribe struct {
	NovelSubscribeId int64      `json:"novelSubscribeId" xorm:"novel_subscribe_id pk autoincr BIGINT"`
	UserId           int64      `json:"userId" xorm:"user_id"`
	NovelId          int64      `json:"novelId" xorm:"novel_id"`
	IsUpgrade        int8       `json:"isUpgrade" xorm:"is_upgrade"`
	CreatedAt        types.Time `json:"createdAt" xorm:"created_at timestampz created"`
	UpdatedAt        types.Time `json:"updatedAt" xorm:"updated_at timestampz updated"`
}

var _ model.BaseModel = &NovelSubscribe{}

func (e *NovelSubscribe) TableName() string {
	return "novel_subscribe"
}
