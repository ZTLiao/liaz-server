package model

import (
	"core/model"
	"core/types"
)

type RecommendItem struct {
	RecommendItemId int64      `json:"recommendItemId" xorm:"recommend_item_id pk autoincr BIGINT"`
	RecommendId     int64      `json:"recommendId" xorm:"recommend_id"`
	Title           string     `json:"title" xorm:"title"`
	SubTitle        string     `json:"subTitle" xorm:"sub_title"`
	ShowValue       string     `json:"showValue" xorm:"show_value"`
	SkipType        int8       `json:"skipType" xorm:"skip_type"`
	SkipValue       string     `json:"skipValue" xorm:"skip_value"`
	ObjId           string     `json:"objId" xorm:"obj_id"`
	SeqNo           int        `json:"seqNo" xorm:"seq_no"`
	Status          int8       `json:"status" xorm:"status"`
	CreatedAt       types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt       types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &RecommendItem{}

func (e *RecommendItem) TableName() string {
	return "recommend_item"
}
