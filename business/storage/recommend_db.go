package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"time"

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
	var now = types.Time(time.Now())
	recommendId := recommend.RecommendId
	if recommendId == 0 {
		recommend.CreatedAt = now
		_, err := e.db.Insert(recommend)
		if err != nil {
			return err
		}
	} else {
		recommend.UpdatedAt = now
		_, err := e.db.ID(recommendId).Cols("title", "position", "recommend_type", "show_type", "show_title", "opt_type", "opt_value", "seq_no", "status", "updated_at").Update(recommend)
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

func (e *RecommendDb) GetRecommendByPosition(position int8) ([]model.Recommend, error) {
	var recommends []model.Recommend
	err := e.db.Where("position = ? and status = ?", position, constant.PASS).OrderBy("seq_no asc").Find(&recommends)
	if err != nil {
		return nil, err
	}
	return recommends, nil
}

func (e *RecommendDb) GetRecommendById(recommendId int64) (*model.Recommend, error) {
	if recommendId == 0 {
		return nil, nil
	}
	var recommend model.Recommend
	_, err := e.db.ID(recommendId).Get(&recommend)
	if err != nil {
		return nil, err
	}
	if recommend.RecommendId == 0 {
		return nil, nil
	}
	return &recommend, nil
}
