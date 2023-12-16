package model

import (
	"core/model"
	"core/types"
)

type Comic struct {
	ComicId      int64      `json:"comicId" xorm:"comic_id pk autoincr BIGINT"`
	Title        string     `json:"title" xorm:"title"`
	Cover        string     `json:"cover" xorm:"cover"`
	Description  string     `json:"description" xorm:"description"`
	FirstLetter  string     `json:"firstLetter" xorm:"first_letter"`
	Flag         int8       `json:"flag" xorm:"flag"`
	CategoryIds  string     `json:"categoryIds" xorm:"category_ids"`
	Categorys    string     `json:"categorys" xorm:"categorys"`
	AuthorIds    string     `json:"authorIds" xorm:"author_ids"`
	Authors      string     `json:"authors" xorm:"authors"`
	ChapterNum   int32      `json:"chapterNum" xorm:"chapter_num"`
	StartTime    types.Time `json:"startTime" xorm:"start_time"`
	EndTime      types.Time `json:"endTime" xorm:"end_time"`
	SubscribeNum int32      `json:"subscribeNum" xorm:"subscribe_num"`
	HitNum       int32      `json:"hitNum" xorm:"hit_num"`
	Status       int8       `json:"status" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt    types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &Comic{}

func (e *Comic) TableName() string {
	return "comic"
}
