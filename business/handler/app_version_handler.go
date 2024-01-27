package handler

import (
	"basic/device"
	"business/storage"
	"core/response"
	"core/utils"
	"core/web"
)

type AppVersionHandler struct {
	AppVersionDb *storage.AppVersionDb
}

func (e *AppVersionHandler) CheckUpdate(wc *web.WebContext) interface{} {
	deviceInfo := device.GetDeviceInfo(wc)
	os := deviceInfo.Os
	if len(os) == 0 {
		return response.Success()
	}
	channel := deviceInfo.Channel
	if len(channel) == 0 {
		return response.Success()
	}
	appVersion1 := deviceInfo.AppVersion
	if len(appVersion1) == 0 {
		return response.Success()
	}
	appVersion, err := e.AppVersionDb.GetLatest(os, channel)
	if err != nil {
		wc.AbortWithError(err)
	}
	if appVersion == nil {
		return response.Success()
	}
	appVersion2 := appVersion.AppVersion
	res, err := utils.CompareAppVersion(appVersion1, appVersion2)
	if err != nil {
		wc.AbortWithError(err)
	}
	if res >= 0 {
		return response.Success()
	}
	return response.ReturnOK(appVersion)
}
