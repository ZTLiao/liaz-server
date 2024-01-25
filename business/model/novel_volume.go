package model

import (
	"core/model"
	"core/types"
)

type NovelVolume struct {
	NovelVolumeId int64      `json:"novelVolumeId" xorm:"novel_volume_id pk autoincr BIGINT"`
	NovelId       int64      `json:"novelId" xorm:"novel_id"`
	VolumeName    string     `json:"volumeName" xorm:"volume_name"`
	SeqNo         int64      `json:"seqNo" xorm:"seq_no"`
	Status        int8       `json:"status" xorm:"status"`
	CreatedAt     types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt     types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &NovelVolume{}

func (e *NovelVolume) TableName() string {
	return "novel_volume"
}
