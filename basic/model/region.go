package model

import (
	"core/model"
	"core/types"
)

type Region struct {
	RegionId   int64      `json:"regionId" xorm:"region_id pk autoincr BIGINT"`
	RegionName string     `json:"regionName" xorm:"region_name"`
	SeqNo      int32      `json:"seqNo" xorm:"seq_no"`
	Status     int8       `json:"status" xorm:"status"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt  types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &Region{}

func (e *Region) TableName() string {
	return "region"
}
