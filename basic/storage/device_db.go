package storage

import (
	"basic/model"

	"github.com/go-xorm/xorm"
)

type DeviceDb struct {
	db *xorm.Engine
}

func NewDeviceDb(db *xorm.Engine) *DeviceDb {
	return &DeviceDb{db}
}

func (e *DeviceDb) SaveOrUpdateDevice(device *model.Device) error {
	deviceId := device.DeviceId
	count, err := e.db.Where("device_id = ?", deviceId).Count(device)
	if err != nil {
		return err
	}
	if count == 0 {
		_, err := e.db.Insert(device)
		if err != nil {
			return err
		}
	} else {
		_, err := e.db.ID(deviceId).Update(device)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *DeviceDb) IsUpgrade(deviceId string, app string, appVersion string) (bool, error) {
	count, err := e.db.Where("device_id = ? and app = ? and app_version = ?", deviceId, app, appVersion).Count(&model.Device{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
