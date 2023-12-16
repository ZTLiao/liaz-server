package storage

import "github.com/go-xorm/xorm"

type ComicChapterDb struct {
	db *xorm.Engine
}

func NewComicChapterDb(db *xorm.Engine) *ComicChapterDb {
	return &ComicChapterDb{db}
}
