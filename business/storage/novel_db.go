package storage

import "github.com/go-xorm/xorm"

type NovelDb struct {
	db *xorm.Engine
}

func NewNovelDb(db *xorm.Engine) *NovelDb {
	return &NovelDb{db}
}
