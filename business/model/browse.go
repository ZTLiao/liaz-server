package model

import (
	"core/model"
	"core/types"
)

type Browse struct {
	BrowseId    int64      `json:"browseId" xorm:"browse_id pk autoincr BIGINT"`
	UserId      int64      `json:"userId" xorm:"user_id"`
	ObjId       int64      `json:"objId" xorm:"obj_id"`
	AssetType   int8       `json:"assetType" xorm:"asset_type"`
	Title       string     `json:"title" xorm:"title"`
	Cover       string     `json:"cover" xorm:"cover"`
	ChapterId   int64      `json:"chapterId" xorm:"chapter_id"`
	ChapterName string     `json:"chapterName" xorm:"chapter_name"`
	Path        string     `json:"path" xorm:"path"`
	StopIndex   int        `json:"stopIndex" xorm:"stop_index"`
	CreatedAt   types.Time `json:"createdAt" xorm:"created_at timestampz created"`
	UpdatedAt   types.Time `json:"updatedAt" xorm:"updated_at timestampz updated"`
}

var _ model.BaseModel = &Browse{}

func (e *Browse) TableName() string {
	return "browse"
}
