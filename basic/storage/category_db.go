package storage

import (
	"basic/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type CategoryDb struct {
	db *xorm.Engine
}

func NewCategoryDb(db *xorm.Engine) *CategoryDb {
	return &CategoryDb{db}
}

func (e *CategoryDb) GetCategoryPage(startRow int, endRow int) ([]model.Category, int64, error) {
	var categorys []model.Category
	err := e.db.OrderBy("seq_no asc").Limit(endRow, startRow).Find(&categorys)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.Category{})
	if err != nil {
		return nil, 0, err
	}
	return categorys, total, nil
}

func (e *CategoryDb) GetCategoryList() ([]model.Category, error) {
	var categorys []model.Category
	err := e.db.OrderBy("seq_no asc").Find(&categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (e *CategoryDb) SaveOrUpdateCategory(category *model.Category) error {
	var now = types.Time(time.Now())
	categoryId := category.CategoryId
	if categoryId == 0 {
		category.UpdatedAt = now
		_, err := e.db.Insert(category)
		if err != nil {
			return err
		}
	} else {
		category.UpdatedAt = now
		_, err := e.db.ID(categoryId).Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *CategoryDb) DelCategory(categoryId int64) error {
	_, err := e.db.ID(categoryId).Delete(&model.Category{})
	if err != nil {
		return err
	}
	return nil
}

func (e *CategoryDb) GetValidCategory() ([]model.Category, error) {
	var categories []model.Category
	err := e.db.Where("status = ?", constant.YES).OrderBy("seq_no asc").Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (e *CategoryDb) GetCategoryByGroupCode(groupCode string) ([]model.Category, error) {
	var categories []model.Category
	err := e.db.Where("status = ? and group_code = ?", constant.YES, groupCode).Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
