package model

import (
	"core/model"
	"core/types"
)

type ComicSubscribe struct {
	ComicSubscribeId int64      `json:"comicSubscribeId" xorm:"comic_subscribe_id pk autoincr BIGINT"`
	UserId           int64      `json:"userId" xorm:"user_id"`
	ComicId          int64      `json:"comicId" xorm:"comic_id"`
	IsUpgrade        int8       `json:"isUpgrade" xorm:"is_upgrade"`
	CreatedAt        types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt        types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &ComicSubscribe{}

func (e *ComicSubscribe) TableName() string {
	return "comic_subscribe"
}
