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

type NovelDb struct {
	db *xorm.Engine
}

func NewNovelDb(db *xorm.Engine) *NovelDb {
	return &NovelDb{db}
}

func (e *NovelDb) GetNovelUpgradePage(pageNum int32, pageSize int32) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("status = ?", constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *NovelDb) GetNovelById(novelId int64) (*model.Novel, error) {
	var novel model.Novel
	_, err := e.db.Where("novel_id = ? and status = ?", novelId, constant.PASS).Get(&novel)
	if err != nil {
		return nil, err
	}
	return &novel, nil
}

func (e *NovelDb) GetNovelByAuthor(authorId int64, pageNum int32, pageSize int32) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("find_in_set(?, author_ids) and status = ?", authorId, constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *NovelDb) GetNovelByCategory(categoryId int64, pageNum int32, pageSize int32) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("find_in_set(?, category_ids) and status = ?", categoryId, constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *NovelDb) UpdateHitNum(novelId int64, hitNum int32) error {
	if novelId == 0 {
		return nil
	}
	_, err := e.db.ID(novelId).Cols("hit_num", "updated_at").Update(&model.Novel{
		HitNum:    hitNum,
		UpdatedAt: types.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelDb) UpdateSubscribeNum(novelId int64, subscribeNum int32) error {
	if novelId == 0 {
		return nil
	}
	_, err := e.db.ID(novelId).Cols("subscribe_num", "updated_at").Update(&model.Novel{
		SubscribeNum: subscribeNum,
		UpdatedAt:    types.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *NovelDb) GetNovelByAuthorId(authorId int64) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("find_in_set(?, author_ids) and status = ?", authorId, constant.PASS).OrderBy("end_time desc").Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *NovelDb) GetNovelByCategoryId(categoryId int64) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("find_in_set(?, category_ids) and status = ?", categoryId, constant.PASS).OrderBy("end_time desc").Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *NovelDb) GetNovelMapByIds(novelIds []int64) (map[int64]model.Novel, error) {
	if len(novelIds) == 0 {
		return nil, nil
	}
	var novels []model.Novel
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString("novel_id in (")
	for i, length := 0, len(novelIds); i < length; i++ {
		builder.WriteString(utils.QUESTION)
		params = append(params, novelIds[i])
		if i != length-1 {
			builder.WriteString(utils.COMMA)
		}
	}
	builder.WriteString(")")
	err := e.db.Where(builder.String(), params...).Find(&novels)
	if err != nil {
		return nil, err
	}
	if len(novels) == 0 {
		return nil, nil
	}
	var novelMap = make(map[int64]model.Novel, 0)
	for _, v := range novels {
		novelMap[v.NovelId] = v
	}
	return novelMap, nil
}
