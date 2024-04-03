package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"core/utils"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

type ComicDb struct {
	db *xorm.Engine
}

func NewComicDb(db *xorm.Engine) *ComicDb {
	return &ComicDb{db}
}

func (e *ComicDb) GetComicUpgradePage(pageNum int32, pageSize int32) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("status = ?", constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) GetComicById(comicId int64) (*model.Comic, error) {
	var comic model.Comic
	_, err := e.db.Where("comic_id = ? and status = ?", comicId, constant.PASS).Get(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func (e *ComicDb) GetComicByAuthor(authorId int64, pageNum int32, pageSize int32) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("find_in_set(?, author_ids) and status = ?", authorId, constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) GetComicByCategory(categoryId int64, pageNum int32, pageSize int32) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("find_in_set(?, category_ids) and status = ?", categoryId, constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) UpdateHitNum(comicId int64, hitNum int32) error {
	if comicId == 0 {
		return nil
	}
	_, err := e.db.ID(comicId).Cols("hit_num", "updated_at").Update(&model.Comic{
		HitNum:    hitNum,
		UpdatedAt: types.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *ComicDb) UpdateSubscribeNum(comicId int64, subscribeNum int32) error {
	if comicId == 0 {
		return nil
	}
	_, err := e.db.ID(comicId).Cols("subscribe_num", "updated_at").Update(&model.Comic{
		SubscribeNum: subscribeNum,
		UpdatedAt:    types.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *ComicDb) GetComicByAuthorId(authorId int64) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("find_in_set(?, author_ids) and status = ?", authorId, constant.PASS).OrderBy("end_time desc").Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) GetComicByCategoryId(categoryId int64) ([]model.Comic, error) {
	var comics []model.Comic
	err := e.db.Where("find_in_set(?, category_ids) and status = ?", categoryId, constant.PASS).OrderBy("end_time desc").Find(&comics)
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (e *ComicDb) GetComicMapByIds(comicIds []int64) (map[int64]model.Comic, error) {
	if len(comicIds) == 0 {
		return nil, nil
	}
	var comics []model.Comic
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString("comic_id in (")
	for i, length := 0, len(comicIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, comicIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(")")
	err := e.db.Where(builder.String(), params...).Find(&comics)
	if err != nil {
		return nil, err
	}
	if len(comics) == 0 {
		return nil, nil
	}
	var comicMap = make(map[int64]model.Comic, 0)
	for _, v := range comics {
		if v.ComicId != 0 {
			comicMap[v.ComicId] = v
		}
	}
	return comicMap, nil
}

func (e *ComicDb) GetComicPage(searchKey string, startRow int, endRow int) ([]model.Comic, int64, error) {
	var comics []model.Comic
	session := e.db.OrderBy("updated_at desc")
	if len(searchKey) != 0 {
		session = session.And("(title = ? or categories = ? or authors = ?)", searchKey, searchKey, searchKey)
	}
	err := session.Limit(endRow, startRow).Find(&comics)
	if err != nil {
		return nil, 0, err
	}
	var total int64
	if len(searchKey) != 0 {
		total, err = e.db.Where("(title = ? or categories = ? or authors = ?)", searchKey, searchKey, searchKey).Count(&model.Comic{})
	} else {
		total, err = e.db.Count(&model.Comic{})
	}
	if err != nil {
		return nil, 0, err
	}
	return comics, total, nil
}
