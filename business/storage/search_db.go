package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type SearchDb struct {
	db *xorm.Engine
}

func NewSearchDb(db *xorm.Engine) *SearchDb {
	return &SearchDb{db}
}

func (e *SearchDb) SaveSearch(searchKey string, deviceId string, userId int64, result string) error {
	_, err := e.db.Insert(&model.Search{
		SearchKey: searchKey,
		DeviceId:  deviceId,
		UserId:    userId,
		Result:    result,
		Status:    constant.YES,
		CreatedAt: types.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	return nil
}
