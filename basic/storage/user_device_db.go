package storage

import (
	"basic/model"
	"core/constant"
	"core/logger"
	"core/types"
	"time"

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
	var now = types.Time(time.Now())
	var userDevice = new(model.UserDevice)
	userDevice.UserId = userId
	userDevice.DeviceId = deviceId
	userDevice.IsUsed = constant.NO
	userDevice.UpdatedAt = now
	_, err := e.db.Where("user_id = ? and device_id != ?", userId, deviceId).Cols("is_used", "updated_at").Update(userDevice)
	if err != nil {
		return err
	}
	count, err := e.db.Where("user_id = ? and device_id = ?", userId, deviceId).Count(&model.UserDevice{})
	if err != nil {
		return err
	}
	logger.Info("userId : %v, deviceId : %v, count : %v", userId, deviceId, count)
	userDevice.IsUsed = constant.YES
	if count == 0 {
		userDevice.CreatedAt = now
		_, err := e.db.Insert(userDevice)
		if err != nil {
			return err
		}
	} else {
		userDevice.UpdatedAt = now
		_, err := e.db.Update(userDevice)
		if err != nil {
			return err
		}
	}
	return nil
}
