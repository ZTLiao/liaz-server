package handler

import (
	"basic/model"
	"basic/storage"
	"core/constant"
	"core/errors"
	"core/file"
	"core/redis"
	"core/response"
	"core/utils"
	"core/web"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
)

type FileItemHandler struct {
	fileItemDb   *storage.FileItemDb
	fileTemplate file.FileTemplate
}

func NewFileItemHandler(fileItemDb *storage.FileItemDb, fileTemplate file.FileTemplate) *FileItemHandler {
	return &FileItemHandler{
		fileItemDb:   fileItemDb,
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
	fileItem, err := e.Upload(bucketName, fileName, header.Size, data)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(fileItem)
}

func (e *FileItemHandler) Upload(bucketName string, fileName string, fileSize int64, data []byte) (*model.FileItem, error) {
	//获取后缀
	var suffix string
	if strings.Contains(fileName, utils.DOT) {
		suffix = strings.Split(fileName, utils.DOT)[1]
	}
	//判断文件类型
	kind, err := filetype.Match(data)
	if err != nil {
		return nil, err
	}
	fileType := kind.MIME.Value
	if len(fileType) == 0 {
		fileType = constant.TEXT_PLAIN
	}
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	//加锁
	var redisLock = redis.NewRedisLock(timestamp)
	if !redisLock.Lock() {
		return nil, errors.New(http.StatusInternalServerError, constant.UPLOAD_ERROR)
	}
	fileInfo, err := e.fileTemplate.PutObject(bucketName, timestamp, data)
	if err != nil {
		return nil, err
	}
	if fileInfo == nil {
		redisLock.Unlock()
		return nil, errors.New(http.StatusInternalServerError, constant.UPLOAD_ERROR)
	}
	var fileItem = model.FileItem{}
	fileItem.BucketName = bucketName
	fileItem.ObjectName = fileName
	fileItem.Size = fileSize
	fileItem.Path = utils.SLASH + bucketName + utils.SLASH + timestamp
	fileItem.UnqiueId = timestamp
	fileItem.Suffix = suffix
	fileItem.FileType = fileType
	e.fileItemDb.SaveFileItem(&fileItem)
	redisLock.Unlock()
	return &fileItem, nil
}
