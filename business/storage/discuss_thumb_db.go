package storage

import (
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type DiscussThumbDb struct {
	db *xorm.Engine
}

func NewDiscussThumbDb(db *xorm.Engine) *DiscussThumbDb {
	return &DiscussThumbDb{db}
}

func (e *DiscussThumbDb) Save(discussId int64, userId int64) error {
	var now = types.Time(time.Now())
	var discussThumb = &model.DiscussThumb{
		DiscussId: discussId,
		UserId:    userId,
		CreatedAt: now,
	}
	_, err := e.db.Insert(discussThumb)
	if err != nil {
		return err
	}
	return nil
}

func (e *DiscussThumbDb) DelDiscussThumb(discussId int64, userId int64) error {
	_, err := e.db.Where("discuss_id = ? and user_id = ?", discussId, userId).Delete(&model.DiscussThumb{})
	if err != nil {
		return err
	}
	return nil
}
