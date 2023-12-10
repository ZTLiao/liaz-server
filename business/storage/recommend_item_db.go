package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type RecommendItemDb struct {
	db *xorm.Engine
}

func NewRecommendItemDb(db *xorm.Engine) *RecommendItemDb {
	return &RecommendItemDb{db}
}

func (e *RecommendItemDb) GetRecommendItemPage(recommendId int64, startRow int, endRow int) ([]model.RecommendItem, int64, error) {
	var recommendItems []model.RecommendItem
	err := e.db.OrderBy("seq_no asc").Limit(endRow, startRow).Find(&recommendItems)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.RecommendItem{})
	if err != nil {
		return nil, 0, err
	}
	return recommendItems, total, nil
}

func (e *RecommendItemDb) SaveOrUpdateRecommendItem(recommendItem *model.RecommendItem) error {
	recommendItemId := recommendItem.RecommendItemId
	if recommendItemId == 0 {
		_, err := e.db.Insert(recommendItem)
		if err != nil {
			return err
		}
	} else {
		_, err := e.db.ID(recommendItemId).Update(recommendItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *RecommendItemDb) DelRecommendItem(recommendItemId int64) error {
	_, err := e.db.ID(recommendItemId).Delete(&model.RecommendItem{})
	if err != nil {
		return err
	}
	return nil
}
