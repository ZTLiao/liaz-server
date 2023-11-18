package file

import (
	"time"

	"github.com/minio/minio-go/v7"
)

type FileTemplate interface {
	CreateBucket(string) error
	ListObjects(string) []FileObjectInfo
	PutObject(string, string, []byte) (*FileObjectInfo, error)
}

type FileObjectInfo struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastModified"`
	Size         int64     `json:"size"`
	Expires      time.Time `json:"expires"`
	ContentType  string    `json:"contentType"`
}

func NewFileTemplate(client interface{}) FileTemplate {
	var fileTemplate FileTemplate = NewMinioTemplate(client.(*minio.Client))
	return fileTemplate
}
