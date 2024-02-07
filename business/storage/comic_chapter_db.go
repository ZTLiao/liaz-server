package storage

import (
	"business/enums"
	"business/model"
	"business/transfer"
	"core/constant"
	"core/utils"
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
	_, err := e.db.Where("comic_id = ? and status = ?", comicId, constant.PASS).OrderBy("created_at desc, seq_no desc").Limit(1, 0).Get(&comicChapter)
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

func (e *ComicChapterDb) GetUpgradeChapters(comicIds []int64) (map[int64]model.ComicChapter, error) {
	var comicChapters []model.ComicChapter
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString(
		`
		select 
			cc1.comic_id,
			cc1.comic_chapter_id,
			cc1.chapter_name 
		from comic_chapter as cc1 inner join (
			select 
				cc.comic_id, 
				max(cc.comic_chapter_id) as comic_chapter_id 
			from comic_chapter as cc 
			where cc.status = ?
		`)
	params = append(params, constant.PASS)
	builder.WriteString("and cc.comic_id in (")
	for i, length := 0, len(comicIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, comicIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(") group by cc.comic_id")
	builder.WriteString(`
			) as cc2
		on cc1.comic_id = cc2.comic_id
		and cc1.comic_chapter_id = cc2.comic_chapter_id
		order by cc1.created_at desc, cc1.seq_no desc
		`)
	err := e.db.Where(builder.String(), params...).Find(&comicChapters)
	if err != nil {
		return nil, err
	}
	if len(comicChapters) == 0 {
		return nil, nil
	}
	var comicChapterMap = make(map[int64]model.ComicChapter, 0)
	for _, v := range comicChapters {
		if v.ComicId != 0 {
			comicChapterMap[v.ComicId] = v
		}
	}
	return comicChapterMap, nil
}
