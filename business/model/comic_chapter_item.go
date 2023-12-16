package model

import (
	"core/model"
	"core/types"
)

type ComicChapterItem struct {
	ComicChapterItemId int64      `json:"comicChapterItemId" xorm:"comic_chapter_item_id pk autoincr BIGINT"`
	ComicChapterId     int64      `json:"comicChapterId" xorm:"comic_chapter_id"`
	ComicId            int64      `json:"comicId" xorm:"comic_id"`
	Path               string     `json:"path" xorm:"path"`
	SeqNo              int32      `json:"seqNo" xorm:"seq_no"`
	CreatedAt          types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt          types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &ComicChapterItem{}

func (e *ComicChapterItem) TableName() string {
	return "comic_chapter_item"
}
