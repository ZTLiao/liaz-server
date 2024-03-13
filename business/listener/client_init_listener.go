package listener

import (
	"basic/device"
	"basic/model"
	"basic/storage"
	"core/event"
	"core/logger"
)

type ClientInitListener struct {
	deviceDb           *storage.DeviceDb
	clientInitRecordDb *storage.ClientInitRecordDb
}

var _ event.Listener = &ClientInitListener{}

func NewClientInitListener(deviceDb *storage.DeviceDb, clientInitRecordDb *storage.ClientInitRecordDb) *ClientInitListener {
	return &ClientInitListener{
		deviceDb:           deviceDb,
		clientInitRecordDb: clientInitRecordDb,
	}
}

func (e *ClientInitListener) OnListen(event event.Event) {
	source := event.Source
	if source == nil {
		return
	}
	device := source.(*device.DeviceInfo)
	logger.Info("deviceInfo : %s", device)
	deviceId := device.DeviceId
	if len(deviceId) == 0 {
		return
	}
	//是否更新设备
	err := e.deviceDb.SaveOrUpdateDevice(&model.Device{
		DeviceId:   device.DeviceId,
		Os:         device.Os,
		OsVersion:  device.OsVersion,
		App:        device.App,
		AppVersion: device.AppVersion,
		Model:      device.Model,
		Imei:       device.Imei,
		Channel:    device.Channel,
	})
	if err != nil {
		logger.Error(err.Error())
	}
	//APP初始化记录
	e.clientInitRecordDb.InsertClientInitRecord(device)
}
