package storage

import (
	"basic/model"
	"core/constant"

	"github.com/go-xorm/xorm"
)

type UserDeviceDb struct {
	db *xorm.Engine
}

func NewUserDeviceDb(db *xorm.Engine) *UserDeviceDb {
	return &UserDeviceDb{db}
}

func (e *UserDeviceDb) SaveOrUpdateUserDevice(userId int64, deviceId string) error {
	if userId == 0 || len(deviceId) == 0 {
		return nil
	}
	var userDevice = new(model.UserDevice)
	userDevice.UserId = userId
	userDevice.DeviceId = deviceId
	userDevice.IsUsed = constant.NO
	e.db.Where("user_id = ? and device_id != ?", userId, deviceId).Cols("is_used").Update(userDevice)
	count, err := e.db.Where("user_id = ? and device_id = ?", userId, deviceId).Count(userDevice)
	if err != nil {
		return err
	}
	userDevice.IsUsed = constant.YES
	if count == 0 {
		_, err := e.db.Insert(userDevice)
		if err != nil {
			return err
		}
	} else {
		_, err := e.db.Update(userDevice)
		if err != nil {
			return err
		}
	}
	return nil
}
