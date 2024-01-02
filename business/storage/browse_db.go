package storage

import (
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type BrowseDb struct {
	db *xorm.Engine
}

func NewBrowseDb(db *xorm.Engine) *BrowseDb {
	return &BrowseDb{db}
}

func (e *BrowseDb) SaveOrUpdateBrowse(userId int64, objId int64, assetType int8, title string, cover string, chapterId int64, chapterName string, path string, stopIndex int) error {
	var now = types.Time(time.Now())
	var browse model.Browse
	ex, err := e.db.Where("user_id = ? and obj_id = ? and asset_type = ?", userId, objId, assetType).Get(&browse)
	if err != nil {
		return err
	}
	if ex {
		browse.Cover = cover
		browse.ChapterId = chapterId
		browse.ChapterName = chapterName
		browse.Path = path
		browse.StopIndex = stopIndex
		browse.UpdatedAt = now
		_, err := e.db.ID(browse.BrowseId).Update(&browse)
		if err != nil {
			return err
		}
	} else {
		browse.UserId = userId
		browse.ObjId = objId
		browse.AssetType = assetType
		browse.Title = title
		browse.Cover = cover
		browse.ChapterId = chapterId
		browse.ChapterName = chapterName
		browse.Path = path
		browse.StopIndex = stopIndex
		browse.CreatedAt = now
		browse.UpdatedAt = now
		_, err := e.db.Insert(&browse)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *BrowseDb) GetBrowseByObjId(userId int64, assetType int, objId int64) (*model.Browse, error) {
	var browse model.Browse
	_, err := e.db.Where("user_id = ? and asset_type = ? and obj_id = ?", userId, assetType, objId).Get(&browse)
	if err != nil {
		return nil, err
	}
	return &browse, nil
}
