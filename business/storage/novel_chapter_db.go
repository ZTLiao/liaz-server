package storage

import "github.com/go-xorm/xorm"

type NovelChapterDb struct {
	db *xorm.Engine
}

func NewNovelChapterDb(db *xorm.Engine) *NovelChapterDb {
	return &NovelChapterDb{db}
}
