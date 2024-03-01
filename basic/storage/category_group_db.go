package storage

import (
	"basic/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type CategoryGroupDb struct {
	db *xorm.Engine
}

func NewCategoryGroupDb(db *xorm.Engine) *CategoryGroupDb {
	return &CategoryGroupDb{db}
}

func (e *CategoryGroupDb) GetCategoryGroupPage(startRow int, endRow int) ([]model.CategoryGroup, int64, error) {
	var categoryGroups []model.CategoryGroup
	err := e.db.OrderBy("seq_no asc").Limit(endRow, startRow).Find(&categoryGroups)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.CategoryGroup{})
	if err != nil {
		return nil, 0, err
	}
	return categoryGroups, total, nil
}

func (e *CategoryGroupDb) GetCategoryGroupList() ([]model.CategoryGroup, error) {
	var categoryGroups []model.CategoryGroup
	err := e.db.OrderBy("seq_no asc").Find(&categoryGroups)
	if err != nil {
		return nil, err
	}
	return categoryGroups, nil
}

func (e *CategoryGroupDb) SaveOrUpdateCategoryGroup(categoryGroup *model.CategoryGroup) error {
	var now = types.Time(time.Now())
	categoryGroupId := categoryGroup.CategoryGroupId
	if categoryGroupId == 0 {
		categoryGroup.UpdatedAt = now
		_, err := e.db.Insert(categoryGroup)
		if err != nil {
			return err
		}
	} else {
		categoryGroup.UpdatedAt = now
		_, err := e.db.ID(categoryGroupId).Update(categoryGroup)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *CategoryGroupDb) DelCategoryGroup(categoryGroupId int64) error {
	_, err := e.db.ID(categoryGroupId).Delete(&model.CategoryGroup{})
	if err != nil {
		return err
	}
	return nil
}

func (e *CategoryGroupDb) GetValidCategoryGroup() ([]model.CategoryGroup, error) {
	var categoryGroups []model.CategoryGroup
	err := e.db.Where("status = ?", constant.YES).OrderBy("seq_no asc").Find(&categoryGroups)
	if err != nil {
		return nil, err
	}
	return categoryGroups, nil
}

func (e *CategoryGroupDb) GetCategoryGroupById(categoryGroupId int64) (*model.CategoryGroup, error) {
	var categoryGroup model.CategoryGroup
	_, err := e.db.ID(categoryGroupId).Get(&categoryGroup)
	if err != nil {
		return nil, err
	}
	return &categoryGroup, nil
}
