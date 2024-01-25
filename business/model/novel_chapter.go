package model

import (
	"core/model"
	"core/types"
)

type NovelChapter struct {
	NovelChapterId int64      `json:"novelChapterId" xorm:"novel_chapter_id pk autoincr BIGINT"`
	NovelVolumeId  int64      `json:"novelVolumeId" xorm:"novel_volume_id"`
	NovelId        int64      `json:"novelId" xorm:"novel_id"`
	ChapterName    string     `json:"chapterName" xorm:"chapter_name"`
	ChapterType    int8       `json:"chapterType" xorm:"chapter_type"`
	PageNum        int32      `json:"pageNum" xorm:"page_num"`
	SeqNo          int64      `json:"seqNo" xorm:"seq_no"`
	Status         int8       `json:"status" xorm:"status"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt      types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &NovelChapter{}

func (e *NovelChapter) TableName() string {
	return "novel_chapter"
}
