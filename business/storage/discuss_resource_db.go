package storage

import (
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type DiscussResourceDb struct {
	db *xorm.Engine
}

func NewDiscussResourceDb(db *xorm.Engine) *DiscussResourceDb {
	return &DiscussResourceDb{db}
}

func (e *DiscussResourceDb) Save(discussId int64, resType int8, path string) error {
	var now = types.Time(time.Now())
	var discussResource = &model.DiscussResource{
		DiscussId: discussId,
		ResType:   resType,
		Path:      path,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err := e.db.Insert(discussResource)
	if err != nil {
		return err
	}
	return nil
}
