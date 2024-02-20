package handler

import (
	"basic/handler"
	"core/constant"
	"core/response"
	"core/web"
	"time"
)

type AdminFileHandler struct {
	FileItemHandler *handler.FileItemHandler
	SysConfHandler  *handler.SysConfHandler
}

func (e *AdminFileHandler) Upload(wc *web.WebContext) interface{} {
	return e.FileItemHandler.UploadFile(wc)
}

func (e *AdminFileHandler) UploadBatch(wc *web.WebContext) interface{} {
	return e.FileItemHandler.UploadBatchFile(wc)
}

func (e *AdminFileHandler) PresignedGetObject(wc *web.WebContext) interface{} {
	bucketName := wc.Param("bucketName")
	objectName := wc.Param("objectName")
	expireTime, err := e.SysConfHandler.GetIntValueByKey(constant.RESOURCE_EXPIRE_TIME)
	if err != nil {
		wc.AbortWithError(err)
	}
	expires := time.Second * 60 * 60 * time.Duration(expireTime)
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
