package storage

import "github.com/go-xorm/xorm"

type ComicChapterItemDb struct {
	db *xorm.Engine
}

func NewComicChapterItemDb(db *xorm.Engine) *ComicChapterItemDb {
	return &ComicChapterItemDb{db}
}
