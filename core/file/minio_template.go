package file

import (
	"bytes"
	"context"
	"core/application"
	"core/logger"

	"github.com/minio/minio-go/v7"
)

type MinioTemplate struct {
}

var _ FileTemplate = &MinioTemplate{}

func NewMinioTemplate() *MinioTemplate {
	return &MinioTemplate{}
}

func (e *MinioTemplate) CreateBucket(bucketName string) {
	var minioClient = application.GetMinioClient()
	ok, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if !ok {
		minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	}
}

func (e *MinioTemplate) ListObjects(bucketName string) []FileObjectInfo {
	var minioClient = application.GetMinioClient()
	var opt = minio.ListObjectsOptions{}
	opt.Set("Content-Type", "image/jpeg")
	var objectInfos = minioClient.ListObjects(context.Background(), bucketName, opt)
	var fileInfos = make([]FileObjectInfo, 0)
	for objectInfo := range objectInfos {
		fileInfos = append(fileInfos, FileObjectInfo{
			Name:         objectInfo.Key,
			ContentType:  objectInfo.ContentType,
			Size:         objectInfo.Size,
			LastModified: objectInfo.LastModified,
			Expires:      objectInfo.Expires,
		})
	}
	return fileInfos
}

func (e *MinioTemplate) PutObject(bucketName string, objectName string, data []byte) *FileObjectInfo {
	e.CreateBucket(bucketName)
	var minioClient = application.GetMinioClient()
	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewBuffer(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &FileObjectInfo{
		Name:         uploadInfo.Key,
		Size:         uploadInfo.Size,
		LastModified: uploadInfo.LastModified,
		Expires:      uploadInfo.Expiration,
	}
}
