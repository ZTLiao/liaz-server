package storage

import (
	"business/model"
	"core/constant"

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

func (e *NovelDb) GetNovelByCategory(categoryId int64, pageNum int32, pageSize int32) ([]model.Novel, error) {
	var novels []model.Novel
	err := e.db.Where("find_in_set(?, category_ids) and status = ?", categoryId, constant.PASS).OrderBy("end_time desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&novels)
	if err != nil {
		return nil, err
	}
	return novels, nil
}
