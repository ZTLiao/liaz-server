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

func (e *AssetDb) GetAssetForUpgrade(limit int64) ([]model.Asset, error) {
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
		limit ?
		`, limit).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) GetAssetForHot(limit int64) ([]model.Asset, error) {
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
		left join comic_chapter as cc on cc.comic_chapter_id = a.obj_id and a.asset_type = 1
		left join novel_chapter as nc on nc.novel_chapter_id = a.obj_id and a.asset_type = 2
		left join comic as c on c.comic_id = cc.comic_id and c.status = 1
		left join novel as n on n.novel_id = nc.novel_id and n.status = 1
		group by a.asset_id
		order by (ifnull(c.hit_num, 0) + ifnull(n.hit_num, 0)) desc 
		limit ?
		`, limit).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) GetAssetForMySubscribe(userId int64, limit int64) ([]model.Asset, error) {
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
		left join comic_chapter as cc on cc.comic_chapter_id = a.obj_id and a.asset_type = 1
		left join novel_chapter as nc on nc.novel_chapter_id = a.obj_id and a.asset_type = 2
		left join comic as c on c.comic_id = cc.comic_id and c.status = 1
		left join novel as n on n.novel_id = nc.novel_id and n.status = 1
		left join comic_subscribe as cs on cs.comic_id = c.comic_id and cs.user_id = ?
		left join novel_subscribe as ns on ns.novel_id = n.novel_id and ns.user_id = ?
		where 
			cs.is_upgrade = 1 
			or ns.is_upgrade = 1
		group by a.asset_id
		order by a.updated_at desc
		limit ?
		`, userId, userId, limit).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
