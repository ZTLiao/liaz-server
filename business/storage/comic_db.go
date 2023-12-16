package storage

import "github.com/go-xorm/xorm"

type ComicDb struct {
	db *xorm.Engine
}

func NewComicDb(db *xorm.Engine) *ComicDb {
	return &ComicDb{db}
}
