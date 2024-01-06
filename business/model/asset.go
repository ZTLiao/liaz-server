package model

import (
	"core/model"
	"core/types"
)

type Asset struct {
	AssetId        int64      `json:"assetId" xorm:"asset_id pk autoincr BIGINT"`
	AssetKey       string     `json:"assetKey" xorm:"asset_key"`
	AssetType      int8       `json:"assetType" xorm:"asset_type"`
	Title          string     `json:"title" xorm:"title"`
	Cover          string     `json:"cover" xorm:"cover"`
	UpgradeChapter string     `json:"upgradeChapter" xorm:"upgrade_chapter"`
	CategoryIds    string     `json:"categoryIds" xorm:"category_ids"`
	AuthorIds      string     `json:"authorIds" xorm:"author_ids"`
	ChapterId      int64      `json:"chapterId"  xorm:"chapter_id"`
	ObjId          int64      `json:"objId" xorm:"obj_id"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestampz created"`
	UpdatedAt      types.Time `json:"updatedAt" xorm:"updated_at timestampz updated"`
}

var _ model.BaseModel = &Asset{}

func (e *Asset) TableName() string {
	return "asset"
}
