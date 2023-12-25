package model

import (
	"core/model"
	"core/types"
)

type Novel struct {
	NovelId      int64      `json:"novelId" xorm:"novel_id pk autoincr BIGINT"`
	Title        string     `json:"title" xorm:"title"`
	Cover        string     `json:"cover" xorm:"cover"`
	Description  string     `json:"description" xorm:"description"`
	FirstLetter  string     `json:"firstLetter" xorm:"first_letter"`
	Flag         int8       `json:"flag" xorm:"flag"`
	Direction    int8       `json:"direction" xorm:"direction"`
	CategoryIds  string     `json:"categoryIds" xorm:"category_ids"`
	Categories   string     `json:"categories" xorm:"categories"`
	AuthorIds    string     `json:"authorIds" xorm:"author_ids"`
	Authors      string     `json:"authors" xorm:"authors"`
	RegionId     int64      `json:"regionId" xorm:"region_id"`
	Region       string     `json:"region" xorm:"region"`
	ChapterNum   int32      `json:"chapterNum" xorm:"chapter_num"`
	StartTime    types.Time `json:"startTime" xorm:"start_time"`
	EndTime      types.Time `json:"endTime" xorm:"end_time"`
	SubscribeNum int32      `json:"subscribeNum" xorm:"subscribe_num"`
	HitNum       int32      `json:"hitNum" xorm:"hit_num"`
	Status       int8       `json:"status" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt    types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &Novel{}

func (e *Novel) TableName() string {
	return "novel"
}
