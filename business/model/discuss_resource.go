package model

import (
	"core/model"
	"core/types"
)

type DiscussResource struct {
	DiscussResourceId int64      `json:"discussResourceId" xorm:"discuss_resource_id pk autoincr BIGINT"`
	DiscussId         int64      `json:"discussId" xorm:"discuss_id"`
	ResType           int8       `json:"resType" xorm:"res_type"`
	Path              string     `json:"path" xorm:"path"`
	SeqNo             int        `json:"seqNo" xorm:"seq_no"`
	CreatedAt         types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt         types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &DiscussResource{}

func (e *DiscussResource) TableName() string {
	return "discuss_resource"
}
