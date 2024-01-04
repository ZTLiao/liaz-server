package storage

import (
	"business/enums"
	"business/model"
	"business/transfer"
	"core/constant"
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
		left join browse as b on b.obj_id = n.novel_id and b.asset_type = 2
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
	err := e.db.SQL(builder.String(), userId).Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&novelChapters)
	if err != nil {
		return nil, err
	}
	return novelChapters, nil
}
