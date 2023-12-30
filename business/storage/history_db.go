package storage

import (
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type HistoryDb struct {
	db *xorm.Engine
}

func NewHistoryDb(db *xorm.Engine) *HistoryDb {
	return &HistoryDb{db}
}

func (e *HistoryDb) SaveHistory(deviceId string, userId int64, objId int64, assetType int8, title string, cover string, chapterId int64, chapterName string, path string, stopIndex int) error {
	var now = types.Time(time.Now())
	var history = new(model.History)
	history.DeviceId = deviceId
	history.UserId = userId
	history.ObjId = objId
	history.AssetType = assetType
	history.Title = title
	history.Cover = cover
	history.ChapterId = chapterId
	history.ChapterName = chapterName
	history.Path = path
	history.StopIndex = stopIndex
	history.CreatedAt = now
	_, err := e.db.Insert(history)
	if err != nil {
		return err
	}
	return nil
}
