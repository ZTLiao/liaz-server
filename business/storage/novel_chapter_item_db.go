package storage

import "github.com/go-xorm/xorm"

type NovelChapterItemDb struct {
	db *xorm.Engine
}

func NewNovelChapterItemDb(db *xorm.Engine) *NovelChapterItemDb {
	return &NovelChapterItemDb{db}
}
