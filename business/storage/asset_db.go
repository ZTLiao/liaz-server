package storage

import "github.com/go-xorm/xorm"

type AssetDb struct {
	db *xorm.Engine
}

func NewAssetDb(db *xorm.Engine) *AssetDb {
	return &AssetDb{db}
}
