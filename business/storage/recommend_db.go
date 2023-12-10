package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type RecommendDb struct {
	db *xorm.Engine
}

func NewRecommendDb(db *xorm.Engine) *RecommendDb {
	return &RecommendDb{db}
}

func (e *RecommendDb) GetRecommendPage(startRow int, endRow int) ([]model.Recommend, int64, error) {
	var recommends []model.Recommend
	err := e.db.OrderBy("seq_no asc").Limit(endRow, startRow).Find(&recommends)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.Recommend{})
	if err != nil {
		return nil, 0, err
	}
	return recommends, total, nil
}

func (e *RecommendDb) GetRecommendList() ([]model.Recommend, error) {
	var recommends []model.Recommend
	err := e.db.OrderBy("seq_no asc").Find(&recommends)
	if err != nil {
		return nil, err
	}
	return recommends, nil
}

func (e *RecommendDb) SaveOrUpdateRecommend(recommend *model.Recommend) error {
	recommendId := recommend.RecommendId
	if recommendId == 0 {
		_, err := e.db.Insert(recommend)
		if err != nil {
			return err
		}
	} else {
		_, err := e.db.ID(recommendId).Update(recommend)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *RecommendDb) DelRecommend(recommendId int64) error {
	_, err := e.db.ID(recommendId).Delete(&model.Recommend{})
	if err != nil {
		return err
	}
	return nil
}
