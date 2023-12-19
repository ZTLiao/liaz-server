package storage

import (
	"business/model"
	"core/constant"

	"github.com/go-xorm/xorm"
)

type ComicDb struct {
	db *xorm.Engine
}

func NewComicDb(db *xorm.Engine) *ComicDb {
	return &ComicDb{db}
}

func (e *ComicDb) GetComicUpgradePage(pageNum int32, pageSize int32) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("status = ?", constant.YES).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) GetComicById(comicId int64) (*model.Comic, error) {
	var comic model.Comic
	_, err := e.db.Where("comic_id = ? and status = ?", comicId, constant.YES).Get(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}
