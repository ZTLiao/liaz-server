package storage

import (
	"business/model"
	"core/constant"

	"github.com/go-xorm/xorm"
)

type NovelChapterDb struct {
	db *xorm.Engine
}

func NewNovelChapterDb(db *xorm.Engine) *NovelChapterDb {
	return &NovelChapterDb{db}
}

func (e *NovelChapterDb) UpgradeChapter(novelId int64) (*model.NovelChapter, error) {
	var novelChapter model.NovelChapter
	_, err := e.db.Where("novel_id = ? and status = ?", novelId, constant.PASS).OrderBy("seq_no desc").Limit(1, 0).Get(&novelChapter)
	if err != nil {
		return nil, err
	}
	return &novelChapter, nil
}

func (e *NovelChapterDb) GetNovelChapters(novelId int64) ([]model.NovelChapter, error) {
	var novelChapters []model.NovelChapter
	err := e.db.Where("novel_id = ? and status = ?", novelId, constant.PASS).OrderBy("seq_no asc").Find(&novelChapters)
	if err != nil {
		return nil, err
	}
	return novelChapters, nil
}

func (e *NovelChapterDb) GetNovelChapterById(novelChapterId int64) (*model.NovelChapter, error) {
	var novelChapter model.NovelChapter
	_, err := e.db.ID(novelChapterId).Get(&novelChapter)
	if err != nil {
		return nil, err
	}
	return &novelChapter, nil
}
