package storage

import (
	"basic/device"
	"basic/model"
	"core/types"
	"core/utils"
	"time"

	"github.com/go-xorm/xorm"
)

type CrashRecordDb struct {
	db *xorm.Engine
}

func NewCrashRecordDb(db *xorm.Engine) *CrashRecordDb {
	return &CrashRecordDb{db}
}

func (e *CrashRecordDb) InsertCrashRecord(err string, stackTrace string, deviceInfo *device.DeviceInfo) error {
	var now = types.Time(time.Now())
	clientIp := deviceInfo.ClientIp
	ipRegion, _ := utils.GetAddress(clientIp)
	var record = &model.CrashRecord{
		Err:        err,
		StackTrace: stackTrace,
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
		CreatedAt:  now,
	}
	e.db.Insert(record)
	return nil
}
