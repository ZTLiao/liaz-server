package storage

import (
	"business/model"
	"business/transfer"
	"core/utils"
	"strings"

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
	err := session.Limit(int(pageSize), int((pageNum-1)*pageSize)).OrderBy("updated_at desc").Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) GetAssetByCategoryIds(assetType int8, categoryIds []int64, pageNum int32, pageSize int32) ([]model.Asset, error) {
	var assets []model.Asset
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString("(")
	for i, length := 0, len(categoryIds); i < length; i++ {
		builder.WriteString("find_in_set(?, category_ids)")
		params = append(params, categoryIds[i])
		if i != length-1 {
			builder.WriteString(utils.SPACE)
			builder.WriteString("or")
			builder.WriteString(utils.SPACE)
		}
	}
	builder.WriteString(")")
	if assetType != 0 {
		builder.WriteString(" and asset_type = ? ")
		params = append(params, assetType)
	}
	err := e.db.Where(builder.String(), params...).Limit(int(pageSize), int((pageNum-1)*pageSize)).OrderBy("updated_at desc").Find(&assets)
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
			a.chapter_id,
			a.obj_id,
			a.created_at,
			a.updated_at 
		from asset as a 
		where 
			a.created_at between date_sub(now(), interval 7 day) and now()
		group by a.obj_id
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
			a.chapter_id,
			a.obj_id,
			a.created_at,
			a.updated_at
		from asset as a
		left join comic_chapter as cc on cc.comic_chapter_id = a.chapter_id and a.asset_type = 1
		left join novel_chapter as nc on nc.novel_chapter_id = a.chapter_id and a.asset_type = 2
		left join comic as c on c.comic_id = cc.comic_id and c.status = 1
		left join novel as n on n.novel_id = nc.novel_id and n.status = 1
		where 
		    (
			c.updated_at between date_sub(now(), interval 30 day) and now()
			or 
			n.updated_at between date_sub(now(), interval 30 day) and now()
			)
		group by a.obj_id
		order by (ifnull(c.hit_num, 0) + ifnull(n.hit_num, 0)) desc, c.updated_at desc, n.updated_at desc 
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
			a.chapter_id,
			a.obj_id,
			a.created_at,
			a.updated_at
		from asset as a
		left join comic_chapter as cc on cc.comic_chapter_id = a.chapter_id and a.asset_type = 1
		left join novel_chapter as nc on nc.novel_chapter_id = a.chapter_id and a.asset_type = 2
		left join comic as c on c.comic_id = cc.comic_id and c.status = 1
		left join novel as n on n.novel_id = nc.novel_id and n.status = 1
		left join comic_subscribe as cs on cs.comic_id = c.comic_id and cs.user_id = ?
		left join novel_subscribe as ns on ns.novel_id = n.novel_id and ns.user_id = ?
		where 
			cs.is_upgrade = 1 
			or ns.is_upgrade = 1
		group by a.obj_id
		order by a.updated_at desc
		limit ?
		`, userId, userId, limit).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) Search(key string, pageNum int32, pageSize int32) ([]transfer.SearchDto, error) {
	var searchs []transfer.SearchDto
	err := e.db.SQL(
		`
		select 
			a.asset_id,
			a.title,
			a.cover,
			a.asset_type,
			group_concat(distinct ca.category_name) as categories,
			group_concat(distinct au.author_name) as authors,
			a.upgrade_chapter,
			a.obj_id
		from asset as a 
		left join category as ca on find_in_set(ca.category_id, a.category_ids)
		left join author as au on find_in_set(au.author_id, a.author_ids)
		where locate(?, a.asset_key) > 0
		group by a.obj_id
		order by a.updated_at desc
		limit ?, ?
		`, key, (pageNum-1)*pageSize, pageSize).Find(&searchs)
	if err != nil {
		return nil, err
	}
	return searchs, nil
}

func (e *AssetDb) GetAssetByIds(assetIds []int64) ([]model.Asset, error) {
	if len(assetIds) == 0 {
		return nil, nil
	}
	var assets []model.Asset
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString("asset_id in (")
	for i, length := 0, len(assetIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, assetIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(")")
	err := e.db.Where(builder.String(), params...).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (e *AssetDb) GetAssetByObjId(objIds []int64, assetType int8) ([]model.Asset, error) {
	if len(objIds) == 0 {
		return nil, nil
	}
	var assets []model.Asset
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString("obj_id in (")
	for i, length := 0, len(objIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, objIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(")")
	if assetType != 0 {
		builder.WriteString(" and asset_type = ? ")
		params = append(params, assetType)
	}
	err := e.db.Where(builder.String(), params...).Find(&assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
