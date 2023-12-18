package storage

import (
	"business/model"
	"core/constant"

	"github.com/go-xorm/xorm"
)

type ComicChapterDb struct {
	db *xorm.Engine
}

func NewComicChapterDb(db *xorm.Engine) *ComicChapterDb {
	return &ComicChapterDb{db}
}

func (e *ComicChapterDb) UpgradeChapter(comicId int64) (*model.ComicChapter, error) {
	var comicChapter model.ComicChapter
	_, err := e.db.Where("comic_id = ? and status = ?", comicId, constant.YES).OrderBy("seq_no desc").Limit(1, 0).Get(&comicChapter)
	if err != nil {
		return nil, err
	}
	return &comicChapter, nil
}
