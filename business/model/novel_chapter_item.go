package model

import (
	"core/model"
	"core/types"
)

type NovelChapterItem struct {
	NovelChapterItemId int64      `json:"novelChapterItemId" xorm:"novel_chapter_item_id pk autoincr BIGINT"`
	NovelChapterId     int64      `json:"novelChapterId" xorm:"novel_chapter_id"`
	NovelId            int64      `json:"novelId" xorm:"novel_id"`
	Path               string     `json:"path" xorm:"path"`
	SeqNo              int32      `json:"seqNo" xorm:"seq_no"`
	CreatedAt          types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt          types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &NovelChapterItem{}

func (e *NovelChapterItem) TableName() string {
	return "novel_chapter_item"
}
