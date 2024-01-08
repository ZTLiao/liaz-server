package file

import (
	"core/constant"
	"core/system"
	"time"
)

type FileTemplate interface {
	CreateBucket(string) error
	ListObjects(string) ([]FileObjectInfo, error)
	PutObject(string, string, []byte) (*FileObjectInfo, error)
	PresignedGetObject(bucketName string, objectName string, headers map[string]string, expires time.Duration) (string, error)
}

type FileObjectInfo struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastModified"`
	Size         int64     `json:"size"`
	Expires      time.Time `json:"expires"`
	ContentType  string    `json:"contentType"`
}

func NewFileTemplate(storage string) FileTemplate {
	var fileTemplate FileTemplate
	if storage == constant.MINIO {
		fileTemplate = NewMinioTemplate(system.GetMinioClient())
	}
	return fileTemplate
}
