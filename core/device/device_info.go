package device

type DeviceInfo struct {
	DeviceId   string `json:"deviceId"`
	Os         string `json:"os"`
	OsVersion  string `json:"osVersion"`
	IspType    string `json:"ispType"`
	NetType    string `json:"netType"`
	App        string `json:"app"`
	AppVersion string `json:"appVersion"`
	Model      string `json:"model"`
	Imei       string `json:"imei"`
	Channel    string `json:"channel"`
	Client     string `json:"client"`
}
