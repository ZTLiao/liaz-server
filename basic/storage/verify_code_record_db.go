package storage

import (
	"basic/device"
	"basic/model"
	"core/types"
	"core/utils"
	"time"

	"github.com/go-xorm/xorm"
)

type VerifyCodeRecordDb struct {
	db *xorm.Engine
}

func NewVerifyCodeRecordDb(db *xorm.Engine) *VerifyCodeRecordDb {
	return &VerifyCodeRecordDb{db}
}

func (e *VerifyCodeRecordDb) InsertVerifyCodeRecord(username string, sendType int8, verifyCode string, resCode string, resMsg string, deviceInfo *device.DeviceInfo) error {
	var now = types.Time(time.Now())
	clientIp := deviceInfo.ClientIp
	ipRegion, _ := utils.GetAddress(clientIp)
	var record = &model.VerifyCodeRecord{
		Username:   username,
		SendType:   sendType,
		VerifyCode: verifyCode,
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
		ResCode:    resCode,
		ResMsg:     resMsg,
		CreatedAt:  now,
	}
	_, err := e.db.Insert(record)
	if err != nil {
		return err
	}
	return nil
}
