package model

import (
	"core/model"
	"core/types"
)

type DiscussThumb struct {
	DiscussThumbId int64      `json:"discussThumbId" xorm:"discuss_thumb_id pk autoincr BIGINT"`
	DiscussId      int64      `json:"discussId" xorm:"discuss_id"`
	UserId         int64      `json:"userId" xorm:"user_id"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &DiscussThumb{}

func (e *DiscussThumb) TableName() string {
	return "discuss_thumb"
}
