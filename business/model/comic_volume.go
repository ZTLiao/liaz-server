package model

import (
	"core/model"
	"core/types"
)

type ComicVolume struct {
	ComicVolumeId int64      `json:"comicVolumeId" xorm:"comic_volume_id pk autoincr BIGINT"`
	ComicId       int64      `json:"comicId" xorm:"comic_id"`
	VolumeName    string     `json:"volumeName" xorm:"volume_name"`
	SeqNo         int64      `json:"seqNo" xorm:"seq_no"`
	Status        int8       `json:"status" xorm:"status"`
	CreatedAt     types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt     types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &ComicVolume{}

func (e *ComicVolume) TableName() string {
	return "comic_volume"
}
