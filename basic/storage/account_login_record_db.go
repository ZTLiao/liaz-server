package storage

import (
	"basic/device"
	"basic/model"
	"core/utils"

	"github.com/go-xorm/xorm"
)

type AccountLoginRecordDb struct {
	db *xorm.Engine
}

func NewAccountLoginRecordDb(db *xorm.Engine) *AccountLoginRecordDb {
	return &AccountLoginRecordDb{db}
}

func (e *AccountLoginRecordDb) InsertAccountLoginRecord(userId int64, deviceInfo *device.DeviceInfo) (*model.AccountLoginRecord, error) {
	clientIp := deviceInfo.ClientIp
	ipRegion, _ := utils.GetAddress(clientIp)
	var record = &model.AccountLoginRecord{
		UserId:     userId,
		DeviceId:   deviceInfo.DeviceId,
		Os:         deviceInfo.Os,
		OsVersion:  deviceInfo.OsVersion,
		App:        deviceInfo.App,
		AppVersion: deviceInfo.AppVersion,
		Model:      deviceInfo.Model,
		Imei:       deviceInfo.Imei,
		Channel:    deviceInfo.Channel,
		IspType:    deviceInfo.IspType,
		NetType:    deviceInfo.NetType,
		ClientIp:   deviceInfo.ClientIp,
		IpRegion:   ipRegion,
	}
	_, err := e.db.Insert(record)
	if err != nil {
		return nil, err
	}
	return record, nil
}
