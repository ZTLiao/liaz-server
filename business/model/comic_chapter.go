package model

import (
	"core/model"
	"core/types"
)

type ComicChapter struct {
	ComicChapterId int64      `json:"comicChapterId" xorm:"comic_chapter_id pk autoincr BIGINT"`
	ComicId        int64      `json:"comicId" xorm:"comic_id"`
	ChapterName    string     `json:"chapterName" xorm:"chapter_name"`
	ChapterType    int8       `json:"chapterType" xorm:"chapter_type"`
	PageNum        int32      `json:"pageNum" xorm:"page_num"`
	SeqNo          int64      `json:"seqNo" xorm:"seq_no"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestampz created"`
	UpdatedAt      types.Time `json:"updatedAt" xorm:"updated_at timestampz updated"`
}

var _ model.BaseModel = &ComicChapter{}

func (e *ComicChapter) TableName() string {
	return "comic_chapter"
}
