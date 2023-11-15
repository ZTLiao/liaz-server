package interfaces

import (
	"basic/application"
	"basic/domain/entity"
	"core/constant"
	"core/file"
	"core/response"
	"core/utils"
	"core/web"
	"io"
	"strconv"
	"strings"
	"time"
)

type FileItem struct {
	fileItemApp  application.FileItemAppInterface
	fileTemplate file.FileTemplate
}

func NewFileItem(fileItemApp application.FileItemAppInterface, fileTemplate file.FileTemplate) *FileItem {
	return &FileItem{
		fileItemApp:  fileItemApp,
		fileTemplate: fileTemplate,
	}
}

func (e *FileItem) UploadFile(wc *web.WebContext) interface{} {
	var bucketName = wc.Context.Param("bucketName")
	if len(bucketName) == 0 {
		return response.Success()
	}
	file, header, err := wc.Context.Request.FormFile("file")
	if err != nil {
		return response.Fail(err.Error())
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return response.Fail(err.Error())
	}
	fileName := header.Filename
	var suffix string
	if strings.Contains(fileName, utils.DOT) {
		suffix = strings.Split(fileName, utils.DOT)[1]
	}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	var fileInfo = e.fileTemplate.PutObject(bucketName, timestamp, data)
	if fileInfo == nil {
		return response.Fail(constant.UPLOAD_ERROR)
	}
	var fileItem = entity.FileItem{}
	fileItem.BucketName = bucketName
	fileItem.ObjectName = fileName
	fileItem.Size = header.Size
	fileItem.Path = bucketName + utils.SLASH + timestamp
	fileItem.UnqiueId = timestamp
	fileItem.Suffix = suffix
	e.fileItemApp.SaveFileItem(&fileItem)
	return response.ReturnOK(fileItem)
}
