package handler

import (
	basicHandler "basic/handler"
	"core/response"
	"core/web"
	"io"
)

type FileHandler struct {
	FileItemHandler *basicHandler.FileItemHandler
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
