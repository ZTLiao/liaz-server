package model

import (
	"core/model"
	"core/types"
)

type History struct {
	HistoryId   int64      `json:"historyId" xorm:"history_id pk autoincr BIGINT"`
	DeviceId    string     `json:"deviceId" xorm:"device_id"`
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
}

var _ model.BaseModel = &History{}

func (e *History) TableName() string {
	return "history"
}
