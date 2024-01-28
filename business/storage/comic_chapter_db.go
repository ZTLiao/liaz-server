package storage

import (
	"business/enums"
	"business/model"
	"business/transfer"
	"core/constant"
	"strconv"
	"strings"

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
	_, err := e.db.Where("comic_id = ? and status = ?", comicId, constant.PASS).OrderBy("seq_no desc").Limit(1, 0).Get(&comicChapter)
	if err != nil {
		return nil, err
	}
	if comicChapter.ComicChapterId == 0 {
		return nil, nil
	}
	return &comicChapter, nil
}

func (e *ComicChapterDb) GetComicChapters(comicId int64) ([]model.ComicChapter, error) {
	var comicChapters []model.ComicChapter
	err := e.db.Where("comic_id = ? and status = ?", comicId, constant.PASS).OrderBy("seq_no desc").Find(&comicChapters)
	if err != nil {
		return nil, err
	}
	return comicChapters, nil
}

func (e *ComicChapterDb) GetComicChapterById(comicChapterId int64) (*model.ComicChapter, error) {
	var comicChapter model.ComicChapter
	_, err := e.db.ID(comicChapterId).Get(&comicChapter)
	if err != nil {
		return nil, err
	}
	if comicChapter.ComicChapterId == 0 {
		return nil, nil
	}
	return &comicChapter, nil
}

func (e *ComicChapterDb) GetBookshelf(userId int64, sortType int32, pageNum int32, pageSize int32) ([]transfer.ComicChapterDto, error) {
	var comicChapters []transfer.ComicChapterDto
	var builder strings.Builder
	builder.WriteString(
		`
		select 
			c.comic_id,
			c.title,
			c.cover,
			c.end_time,
			cs.is_upgrade,
			b.chapter_id,
			b.chapter_name
		from comic_subscribe as cs 
		left join comic as c on c.comic_id = cs.comic_id
		left join browse as b on b.obj_id = c.comic_id and b.user_id = cs.user_id and b.asset_type = 1
		where 
			cs.user_id = ?
		group by c.comic_id
		`)
	if sortType == enums.SORT_TYPE_OF_UPDATE {
		builder.WriteString(" order by c.end_time desc")
	} else if sortType == enums.SORT_TYPE_OF_SUBSCRIBE {
		builder.WriteString(" order by cs.comic_subscribe_id desc")
	} else if sortType == enums.SORT_TYPE_OF_BROWSE {
		builder.WriteString(" order by b.updated_at desc")
	}
	builder.WriteString(" limit " + strconv.FormatInt(int64((pageNum-1)*pageSize), 10) + "," + strconv.FormatInt(int64(pageSize), 10))
	err := e.db.SQL(builder.String(), userId).Find(&comicChapters)
	if err != nil {
		return nil, err
	}
	return comicChapters, nil
}
