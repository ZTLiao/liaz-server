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

func (e *AssetDb) GetAssetByCategoryId(assetType int8, categoryId int64, pageNum int32, pageSize int32) ([]model.Asset, error) {
	var assets []model.Asset
	session := e.db.Where("find_in_set(?, category_ids)", categoryId)
	if assetType != 0 {
		session = session.And("asset_type = ?", assetType)
	}
	err := session.Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) GetAssetForUpgrade() ([]model.Asset, error) {
	var assets []model.Asset
	err := e.db.SQL(
		`
		select
			a.asset_id,
			a.asset_key,
			a.asset_type,
			a.title,
			a.cover,
			a.upgrade_chapter,
			a.category_ids,
			a.author_ids,
			a.obj_id,
			a.created_at,
			a.updated_at 
		from asset as a 
		where 
			a.created_at between date_sub(now(), interval 7 day) and now()
		order by a.updated_at desc 
		limit 9
		`).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
