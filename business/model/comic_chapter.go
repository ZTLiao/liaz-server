package model

import (
	"core/model"
	"core/types"
)

type ComicChapter struct {
	ComicChapterId int64      `json:"comicChapterId" xorm:"comic_chapter_id pk autoincr BIGINT"`
	ComicVolumeId  int64      `json:"comicVolumeId" xorm:"comic_volume_id"`
	ComicId        int64      `json:"comicId" xorm:"comic_id"`
	ChapterName    string     `json:"chapterName" xorm:"chapter_name"`
	ChapterType    int8       `json:"chapterType" xorm:"chapter_type"`
	PageNum        int32      `json:"pageNum" xorm:"page_num"`
	SeqNo          int64      `json:"seqNo" xorm:"seq_no"`
	Status         int8       `json:"status" xorm:"status"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt      types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &ComicChapter{}

func (e *ComicChapter) TableName() string {
	return "comic_chapter"
}
