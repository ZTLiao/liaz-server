package device

import (
	"core/constant"
	"core/web"
)

type DeviceInfo struct {
	DeviceId   string `json:"deviceId"`
	Os         string `json:"os"`
	OsVersion  string `json:"osVersion"`
	IspType    string `json:"ispType"`
	NetType    string `json:"netType" `
	App        string `json:"app"`
	AppVersion string `json:"appVersion"`
	Model      string `json:"model"`
	Imei       string `json:"imei"`
	Channel    string `json:"channel"`
	Client     string `json:"client"`
}

func GetDeviceInfo(wc *web.WebContext) *DeviceInfo {
	var deviceInfo = new(DeviceInfo)
	deviceInfo.DeviceId = wc.GetHeader(constant.X_DEVICE_ID)
	deviceInfo.Os = wc.GetHeader(constant.X_OS)
	deviceInfo.OsVersion = wc.GetHeader(constant.X_OS_VERSION)
	deviceInfo.IspType = wc.GetHeader(constant.X_ISP_TYPE)
	deviceInfo.NetType = wc.GetHeader(constant.X_NET_TYPE)
	deviceInfo.App = wc.GetHeader(constant.X_APP)
	deviceInfo.AppVersion = wc.GetHeader(constant.X_APP_VERSION)
	deviceInfo.Model = wc.GetHeader(constant.X_MODEL)
	deviceInfo.Imei = wc.GetHeader(constant.X_IMEI)
	deviceInfo.Channel = wc.GetHeader(constant.X_CHANNEL)
	deviceInfo.Client = wc.GetHeader(constant.X_CLIENT)
	return deviceInfo
}
