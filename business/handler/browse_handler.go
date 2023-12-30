package handler

import (
	"basic/device"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type BrowseHandler struct {
	BrowseDb  *storage.BrowseDb
	HistoryDb *storage.HistoryDb
}

func (e *BrowseHandler) BrowseHistory(wc *web.WebContext) interface{} {
	objIdStr := wc.PostForm("objId")
	assetTypeStr := wc.PostForm("assetType")
	title := wc.PostForm("title")
	cover := wc.PostForm("cover")
	chapterIdStr := wc.PostForm("chapterId")
	chapterName := wc.PostForm("chapterName")
	path := wc.PostForm("path")
	stopIndexStr := wc.PostForm("stopIndex")
	objId, err := strconv.ParseInt(objIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	assetType, err := strconv.ParseInt(assetTypeStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	chapterId, err := strconv.ParseInt(chapterIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	stopIndex, err := strconv.ParseInt(stopIndexStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	if userId != 0 {
		err := e.BrowseDb.SaveOrUpdateBrowse(userId, objId, int8(assetType), title, cover, chapterId, chapterName, path, int(stopIndex))
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	deviceInfo := device.GetDeviceInfo(wc)
	if deviceInfo != nil {
		err := e.HistoryDb.SaveHistory(deviceInfo.DeviceId, userId, objId, int8(assetType), title, cover, chapterId, chapterName, path, int(stopIndex))
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	return response.Success()
}
