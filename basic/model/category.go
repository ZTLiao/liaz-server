package model

import (
	"core/model"
	"core/types"
)

type Category struct {
	CategoryId   int64      `json:"categoryId" xorm:"category_id"`
	CategoryCode string     `json:"categoryCode" xorm:"category_code"`
	CategoryName string     `json:"categoryName" xorm:"category_name"`
	SeqNo        int32      `json:"seqNo" xorm:"seq_no"`
	Status       int8       `json:"status" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt    types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &Category{}

func (e *Category) TableName() string {
	return "category"
}
