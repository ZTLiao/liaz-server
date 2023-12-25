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
	var err error
	if assetType != 0 {
		err = e.db.Where("asset_type = ? and find_in_set(?, category_ids)", assetType, categoryId).Find(&assets)
	} else {
		err = e.db.Where("find_in_set(?, category_ids)", categoryId).Find(&assets)
	}
	if err != nil {
		return nil, err
	}
	return assets, nil
}
