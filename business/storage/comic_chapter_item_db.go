package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type ComicChapterItemDb struct {
	db *xorm.Engine
}

func NewComicChapterItemDb(db *xorm.Engine) *ComicChapterItemDb {
	return &ComicChapterItemDb{db}
}

func (e *ComicChapterItemDb) GetComicChapterItems(comicId int64) ([]model.ComicChapterItem, error) {
	var comicChapterItems []model.ComicChapterItem
	err := e.db.Where("comic_id = ?", comicId).OrderBy("seq_no asc").Find(&comicChapterItems)
	if err != nil {
		return nil, err
	}
	return comicChapterItems, nil
}
