package listener

import (
	"basic/device"
	"basic/model"
	"basic/storage"
	"core/event"
	"core/logger"
)

type ClientInitListener struct {
	DeviceDb           *storage.DeviceDb
	ClientInitRecordDb *storage.ClientInitRecordDb
}

var _ event.Listener = &ClientInitListener{}

func NewClientInitListener(deviceDb *storage.DeviceDb, ClientInitRecordDb *storage.ClientInitRecordDb) *ClientInitListener {
	return &ClientInitListener{DeviceDb: deviceDb, ClientInitRecordDb: ClientInitRecordDb}
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
	app := device.App
	appVersion := device.AppVersion
	//是否更新设备
	if ok, _ := e.DeviceDb.IsUpgrade(deviceId, app, appVersion); ok {
		err := e.DeviceDb.SaveOrUpdateDevice(&model.Device{
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
			logger.Panic(err.Error())
		}
	}
	//APP初始化记录
	e.ClientInitRecordDb.InsertClientInitRecord(device)
}
