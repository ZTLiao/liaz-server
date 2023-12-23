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

func (e *ComicChapterItemDb) GetComicChapterItemByComicId(comicId int64) ([]model.ComicChapterItem, error) {
	var comicChapterItems []model.ComicChapterItem
	err := e.db.Where("comic_id = ?", comicId).OrderBy("seq_no asc").Find(&comicChapterItems)
	if err != nil {
		return nil, err
	}
	return comicChapterItems, nil
}

func (e *ComicChapterItemDb) GetComicChapterItemByComicChapterId(comicChapterId int64) ([]model.ComicChapterItem, error) {
	var comicChapterItems []model.ComicChapterItem
	err := e.db.Where("comic_chapter_id = ?", comicChapterId).OrderBy("seq_no asc").Find(&comicChapterItems)
	if err != nil {
		return nil, err
	}
	return comicChapterItems, nil
}

func (e *ComicChapterItemDb) GetPathByComicChapterId(comicChapterId int64) ([]string, error) {
	var paths = make([]string, 0)
	comicChapterItems, err := e.GetComicChapterItemByComicChapterId(comicChapterId)
	if err != nil {
		return nil, err
	}
	if len(comicChapterItems) == 0 {
		return paths, nil
	}
	for _, v := range comicChapterItems {
		paths = append(paths, v.Path)
	}
	return paths, nil
}
