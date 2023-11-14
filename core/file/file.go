package file

import (
	"time"
)

type FileTemplate interface {
	CreateBucket(bucketName string)
	ListObjects(bucketName string) []FileObjectInfo
	PutObject(bucketName string, objectName string, data []byte) *FileObjectInfo
}

type FileObjectInfo struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastModified"`
	Size         int64     `json:"size"`
	Expires      time.Time `json:"expires"`
	ContentType  string    `json:"contentType"`
}
