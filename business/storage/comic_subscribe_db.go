package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type ComicSubscribeDb struct {
	db *xorm.Engine
}

func NewComicSubscribeDb(db *xorm.Engine) *ComicSubscribeDb {
	return &ComicSubscribeDb{db}
}

func (e *ComicSubscribeDb) SaveComicSubscribe(comicId int64, userId int64) error {
	if comicId == 0 || userId == 0 {
		return nil
	}
	count, err := e.db.Where("comic_id = ? and user_id = ?", comicId, userId).Count(&model.ComicSubscribe{})
	if err != nil {
		return err
	}
	if count == 0 {
		_, err := e.db.Insert(&model.ComicSubscribe{
			ComicId:   comicId,
			UserId:    userId,
			CreatedAt: types.Time(time.Now()),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ComicSubscribeDb) DelComicSubscribe(comicId int64, userId int64) error {
	if comicId == 0 || userId == 0 {
		return nil
	}
	_, err := e.db.Where("comic_id = ? and user_id = ?", comicId, userId).Delete(&model.ComicSubscribe{})
	if err != nil {
		return err
	}
	return nil
}

func (e *ComicSubscribeDb) IsSubscribe(comicId int64, userId int64) (bool, error) {
	if comicId == 0 || userId == 0 {
		return false, nil
	}
	count, err := e.db.Where("comic_id = ? and user_id = ?", comicId, userId).Count(&model.ComicSubscribe{})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (e *ComicSubscribeDb) SetRead(comicId int64, userId int64) error {
	if comicId == 0 || userId == 0 {
		return nil
	}
	_, err := e.db.Where("comic_id = ? and user_id = ?", comicId, userId).Update(&model.ComicSubscribe{
		IsUpgrade: constant.NO,
	})
	if err != nil {
		return err
	}
	return nil
}
