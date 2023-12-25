package model

import (
	"core/model"
	"core/types"
)

type Asset struct {
	AssetId   int64      `json:"assetId" xorm:"asset_id pk autoincr BIGINT"`
	AssetKey  string     `json:"assetKey" xorm:"asset_key"`
	AssetType int8       `json:"assetType" xorm:"asset_type"`
	ObjId     int64      `json:"objId" xorm:"obj_id"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &Asset{}

func (e *Asset) TableName() string {
	return "asset"
}
