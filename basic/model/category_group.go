package model

import (
	"core/model"
	"core/types"
)

type CategoryGroup struct {
	CategoryGroupId int64      `json:"categoryGroupId" xorm:"category_group_id pk autoincr BIGINT"`
	GroupCode       string     `json:"groupCode" xorm:"group_code"`
	GroupName       string     `json:"groupName" xorm:"group_name"`
	SeqNo           int32      `json:"seqNo" xorm:"seq_no"`
	Status          int8       `json:"status" xorm:"status"`
	CreatedAt       types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt       types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &CategoryGroup{}

func (e *CategoryGroup) TableName() string {
	return "category_group"
}
