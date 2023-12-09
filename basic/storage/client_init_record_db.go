package storage

import (
	"basic/model"

	"github.com/go-xorm/xorm"
)

type ClientInitRecordDb struct {
	db *xorm.Engine
}

func NewClientInitRecordDb(db *xorm.Engine) *ClientInitRecordDb {
	return &ClientInitRecordDb{db}
}

func (e *ClientInitRecordDb) InsertClientInitRecord(clientInitRecord *model.ClientInitRecord) error {
	_, err := e.db.Insert(clientInitRecord)
	if err != nil {
		return err
	}
	return nil
}
