package model

import (
	"core/model"
	"core/types"
)

type Recommend struct {
	RecommendId   int64      `json:"recommendId" xorm:"recommend_id pk autoincr BIGINT"`
	Title         string     `json:"title" xorm:"title"`
	Position      int8       `json:"position" xorm:"position"`
	RecommendType int8       `json:"recommendType" xorm:"recommend_type"`
	ShowType      int8       `json:"showType" xorm:"show_type"`
	ShowTitle     int8       `json:"showTitle" xorm:"show_title"`
	OptType       int8       `json:"optType" xorm:"opt_type"`
	OptValue      string     `json:"optValue" xorm:"opt_value"`
	SeqNo         int        `json:"seqNo" xorm:"seq_no"`
	Status        int8       `json:"status" xorm:"status"`
	CreatedAt     types.Time `json:"createdAt" xorm:"created_at timestampz created"`
	UpdatedAt     types.Time `json:"updatedAt" xorm:"updated_at timestampz updated"`
}

var _ model.BaseModel = &Recommend{}

func (e *Recommend) TableName() string {
	return "recommend"
}
