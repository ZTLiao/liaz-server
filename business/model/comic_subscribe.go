package model

import (
	"core/model"
	"core/types"
)

type ComicSubscribe struct {
	ComicSubscribeId int64      `json:"comicSubscribeId" xorm:"comic_subscribe_id pk autoincr BIGINT"`
	UserId           int64      `json:"userId" xorm:"user_id"`
	ComicId          int64      `json:"comicId" xorm:"comic_id"`
	CreatedAt        types.Time `json:"createdAt" xorm:"created_at timestampz created"`
}

var _ model.BaseModel = &ComicSubscribe{}

func (e *ComicSubscribe) TableName() string {
	return "comic_subscribe"
}
