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
	if novelChapter.NovelChapterId == 0 {
		return nil, nil
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
	if novelChapter.NovelChapterId == 0 {
		return nil, nil
	}
	return &novelChapter, nil
}

func (e *NovelChapterDb) GetBookshelf(userId int64, sortType int32, pageNum int32, pageSize int32) ([]transfer.NovelChapterDto, error) {
	var novelChapters []transfer.NovelChapterDto
	var builder strings.Builder
	builder.WriteString(
		`
		select 
			n.novel_id,
			n.title,
			n.cover,
			n.end_time,
			ns.is_upgrade,
			b.chapter_id,
			b.chapter_name
		from novel_subscribe as ns 
		left join novel as n on n.novel_id = ns.novel_id
		left join browse as b on b.obj_id = n.novel_id and b.user_id = ns.user_id and b.asset_type = 2
		where 
			ns.user_id = ?
		group by n.novel_id
		`)
	if sortType == enums.SORT_TYPE_OF_UPDATE {
		builder.WriteString("order by n.end_time desc")
	} else if sortType == enums.SORT_TYPE_OF_SUBSCRIBE {
		builder.WriteString("order by ns.novel_subscribe_id desc")
	} else if sortType == enums.SORT_TYPE_OF_BROWSE {
		builder.WriteString("order by b.updated_at desc")
	}
	builder.WriteString(" limit " + strconv.FormatInt(int64((pageNum-1)*pageSize), 10) + "," + strconv.FormatInt(int64(pageSize), 10))
	err := e.db.SQL(builder.String(), userId).Find(&novelChapters)
	if err != nil {
		return nil, err
	}
	return novelChapters, nil
}

func (e *NovelChapterDb) GetUpgradeChapters(novelIds []int64) (map[int64]model.NovelChapter, error) {
	var novelChapters []model.NovelChapter
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString(
		`
		select 
			nc1.novel_id,
			nc1.novel_chapter_id,
			nc1.chapter_name 
		from novel_chapter as nc1 inner join (
			select 
				nc.novel_id, 
				max(nc.novel_chapter_id) as novel_chapter_id 
			from novel_chapter as nc 
			where nc.status = ?
		`)
	params = append(params, constant.PASS)
	builder.WriteString("and nc.novel_id in (")
	for i, length := 0, len(novelIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, novelIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(") group by nc.novel_id")
	builder.WriteString(`
			) as nc2
		on nc1.novel_id = nc2.novel_id
		and nc1.novel_chapter_id = nc2.novel_chapter_id
		order by nc1.created_at desc, nc1.seq_no desc
		`)
	err := e.db.Where(builder.String(), params...).Find(&novelChapters)
	if err != nil {
		return nil, err
	}
	if len(novelChapters) == 0 {
		return nil, nil
	}
	var novelChapterMap = make(map[int64]model.NovelChapter, 0)
	for _, v := range novelChapters {
		if v.NovelId != 0 {
			novelChapterMap[v.NovelId] = v
		}
	}
	return novelChapterMap, nil
}
