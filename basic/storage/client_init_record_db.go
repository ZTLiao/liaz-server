package storage

import (
	"basic/device"
	"basic/model"
	"core/types"
	"core/utils"
	"time"

	"github.com/go-xorm/xorm"
)

type ClientInitRecordDb struct {
	db *xorm.Engine
}

func NewClientInitRecordDb(db *xorm.Engine) *ClientInitRecordDb {
	return &ClientInitRecordDb{db}
}

func (e *ClientInitRecordDb) InsertClientInitRecord(deviceInfo *device.DeviceInfo) error {
	var now = types.Time(time.Now())
	clientIp := deviceInfo.ClientIp
	ipRegion, err := utils.GetAddress(clientIp)
	if err != nil {
		return err
	}
	_, err = e.db.Insert(&model.ClientInitRecord{
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
	})
	if err != nil {
		return err
	}
	return nil
}
