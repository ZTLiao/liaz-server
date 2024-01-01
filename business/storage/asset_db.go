package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type AssetDb struct {
	db *xorm.Engine
}

func NewAssetDb(db *xorm.Engine) *AssetDb {
	return &AssetDb{db}
}

func (e *AssetDb) GetAssetByCategoryId(assetType int8, categoryId int64) ([]model.Asset, error) {
	var assets []model.Asset
	session := e.db.Where("find_in_set(?, category_ids)", categoryId)
	if assetType != 0 {
		session = session.And("asset_type = ?", assetType)
	}
	err := session.Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
