package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type NovelSubscribeDb struct {
	db *xorm.Engine
}

func NewNovelSubscribeDb(db *xorm.Engine) *NovelSubscribeDb {
	return &NovelSubscribeDb{db}
}

func (e *NovelSubscribeDb) SaveNovelSubscribe(novelId int64, userId int64) error {
	if novelId == 0 || userId == 0 {
		return nil
	}
	count, err := e.db.Where("novel_id = ? and user_id = ?", novelId, userId).Count(&model.NovelSubscribe{})
	if err != nil {
		return err
	}
	if count == 0 {
		_, err := e.db.Insert(&model.NovelSubscribe{
			NovelId:   novelId,
			UserId:    userId,
			CreatedAt: types.Time(time.Now()),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *NovelSubscribeDb) DelNovelSubscribe(novelId int64, userId int64) error {
	if novelId == 0 || userId == 0 {
		return nil
	}
	_, err := e.db.Where("novel_id = ? and user_id = ?", novelId, userId).Delete(&model.NovelSubscribe{})
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelSubscribeDb) IsSubscribe(novelId int64, userId int64) (bool, error) {
	if novelId == 0 || userId == 0 {
		return false, nil
	}
	count, err := e.db.Where("novel_id = ? and user_id = ?", novelId, userId).Count(&model.NovelSubscribe{})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (e *NovelSubscribeDb) SetRead(novelId int64, userId int64) error {
	if novelId == 0 || userId == 0 {
		return nil
	}
	_, err := e.db.Where("novel_id = ? and user_id = ?", novelId, userId).Update(&model.NovelSubscribe{
		IsUpgrade: constant.NO,
	})
	if err != nil {
		return err
	}
	return nil
}
