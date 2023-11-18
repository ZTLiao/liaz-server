package interfaces

import (
	"basic/application"
	"basic/infrastructure/persistence/entity"
	"core/constant"
	"core/file"
	"core/redis"
	"core/response"
	"core/utils"
	"core/web"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
)

type FileItemHandler struct {
	fileItemApp  application.FileItemAppInterface
	fileTemplate file.FileTemplate
}

func NewFileItemHandler(fileItemApp application.FileItemAppInterface, fileTemplate file.FileTemplate) *FileItemHandler {
	return &FileItemHandler{
		fileItemApp:  fileItemApp,
		fileTemplate: fileTemplate,
	}
}

func (e *FileItemHandler) UploadFile(wc *web.WebContext) interface{} {
	var bucketName = wc.Param("bucketName")
	if len(bucketName) == 0 {
		return response.Success()
	}
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
	//获取后缀
	var suffix string
	if strings.Contains(fileName, utils.DOT) {
		suffix = strings.Split(fileName, utils.DOT)[1]
	}
	//判断文件类型
	kind, err := filetype.Match(data)
	if err != nil {
		wc.AbortWithError(err)
	}
	fileType := kind.MIME.Value
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	//加锁
	var redisLock = redis.NewRedisLock(timestamp)
	if !redisLock.Lock() {
		return response.Fail(constant.UPLOAD_ERROR)
	}
	fileInfo, err := e.fileTemplate.PutObject(bucketName, timestamp, data)
	if fileInfo == nil {
		redisLock.Unlock()
		return response.Fail(constant.UPLOAD_ERROR)
	}
	var fileItem = entity.FileItem{}
	fileItem.BucketName = bucketName
	fileItem.ObjectName = fileName
	fileItem.Size = header.Size
	fileItem.Path = utils.SLASH + bucketName + utils.SLASH + timestamp
	fileItem.UnqiueId = timestamp
	fileItem.Suffix = suffix
	fileItem.FileType = fileType
	e.fileItemApp.SaveFileItem(&fileItem)
	redisLock.Unlock()
	return response.ReturnOK(fileItem)
}
