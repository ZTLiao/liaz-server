package file

import (
	"core/config"
	"core/constant"
	"core/system"
	"time"
)

type FileTemplate interface {
	CreateBucket(string) error
	ListObjects(string) ([]FileObjectInfo, error)
	PutObject(string, string, []byte) (*FileObjectInfo, error)
	PresignedGetObject(string, string, map[string]string, time.Duration) (string, error)
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
	} else if storage == constant.OSS {
		fileTemplate = NewOssTemplate(config.SystemConfig.Oss.BucketName, system.GetOssClient())
	} else if storage == constant.COS {
		fileTemplate = NewCosTemplate(system.GetCosClient())
	}
	return fileTemplate
}
