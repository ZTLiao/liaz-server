package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type NovelChapterItemDb struct {
	db *xorm.Engine
}

func NewNovelChapterItemDb(db *xorm.Engine) *NovelChapterItemDb {
	return &NovelChapterItemDb{db}
}

func (e *NovelChapterItemDb) GetNovelChapterItemByNovelId(novelId int64) ([]model.NovelChapterItem, error) {
	var novelChapterItems []model.NovelChapterItem
	err := e.db.Where("novel_id = ?", novelId).OrderBy("seq_no asc").Find(&novelChapterItems)
	if err != nil {
		return nil, err
	}
	return novelChapterItems, nil
}

func (e *NovelChapterItemDb) GetNovelChapterItemByNovelChapterId(novelChapterId int64) ([]model.NovelChapterItem, error) {
	var novelChapterItems []model.NovelChapterItem
	err := e.db.Where("novel_chapter_id = ?", novelChapterId).OrderBy("seq_no asc").Find(&novelChapterItems)
	if err != nil {
		return nil, err
	}
	return novelChapterItems, nil
}

func (e *NovelChapterItemDb) GetPathByNovelChapterId(novelChapterId int64) ([]string, error) {
	var paths = make([]string, 0)
	novelChapterItems, err := e.GetNovelChapterItemByNovelChapterId(novelChapterId)
	if err != nil {
		return nil, err
	}
	if len(novelChapterItems) == 0 {
		return paths, nil
	}
	for _, v := range novelChapterItems {
		paths = append(paths, v.Path)
	}
	return paths, nil
}
