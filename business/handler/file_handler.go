package handler

import (
	basicHandler "basic/handler"
	"core/constant"
	"core/response"
	"core/web"
	"io"
	"time"
)

type FileHandler struct {
	FileItemHandler *basicHandler.FileItemHandler
	SysConfHandler  *basicHandler.SysConfHandler
}

func (e *FileHandler) Upload(wc *web.WebContext) interface{} {
	bucketName := wc.PostForm("bucketName")
	file, header, err := wc.FormFile("file")
	if err != nil {
		return response.Fail(err.Error())
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return response.Fail(err.Error())
	}
	fileName := header.Filename
	fileSize := header.Size
	wc.Info("bucketName : %s, fileName : %s, fileSize : %d", bucketName, fileName, fileSize)
	fileItem, err := e.FileItemHandler.Upload(bucketName, fileName, fileSize, data)
	if err != nil {
		wc.AbortWithError(err)
	}
	if fileItem == nil {
		return response.Success()
	}
	return response.ReturnOK(fileItem.Path)
}

func (e *FileHandler) PresignedGetObject(wc *web.WebContext) interface{} {
	bucketName := wc.Param("bucketName")
	objectName := wc.Param("objectName")
	expireTime, err := e.SysConfHandler.GetIntValueByKey(constant.RESOURCE_EXPIRE_TIME)
	if err != nil {
		wc.AbortWithError(err)
	}
	expires := time.Second * 24 * 60 * 60 * time.Duration(expireTime)
	if expires == 0 {
		return response.Success()
	}
	var headers = make(map[string]string)
	requestURI, err := e.FileItemHandler.PresignedGetObject(bucketName, objectName, headers, expires)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(requestURI)
}
