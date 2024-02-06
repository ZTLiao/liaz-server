package storage

import (
	"basic/device"
	"business/model"
	"core/types"
	"core/utils"
	"time"

	"github.com/go-xorm/xorm"
)

type AdvertRecordDb struct {
	db *xorm.Engine
}

func NewAdvertRecordDb(db *xorm.Engine) *AdvertRecordDb {
	return &AdvertRecordDb{db}
}

func (e *AdvertRecordDb) InsertAdvertRecord(advertType string, content string, deviceInfo *device.DeviceInfo) error {
	var now = types.Time(time.Now())
	clientIp := deviceInfo.ClientIp
	ipRegion, _ := utils.GetAddress(clientIp)
	var record = &model.AdvertRecord{
		AdvertType: advertType,
		Content:    content,
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
