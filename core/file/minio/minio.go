package minio

import (
	"bytes"
	"context"
	"core/application"
	"core/file"
	"core/logger"

	"github.com/minio/minio-go/v7"
)

type MinioTemplate struct {
}

func (e *MinioTemplate) CreateBucket(bucketName string) {
	var minioClient = application.GetMinioClient()
	ok, err := minioClient.BucketExists(context.Background(), bucketName)
	if err == nil && !ok {
		minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	} else {
		logger.Error(err.Error())
	}
}

func (e *MinioTemplate) ListObjects(bucketName string) []file.FileObjectInfo {
	var minioClient = application.GetMinioClient()
	var objectInfos = minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
	var fileInfos = make([]file.FileObjectInfo, 0)
	for objectInfo := range objectInfos {
		fileInfos = append(fileInfos, file.FileObjectInfo{
			Name:         objectInfo.Key,
			ContentType:  objectInfo.ContentType,
			Size:         objectInfo.Size,
			LastModified: objectInfo.LastModified,
			Expires:      objectInfo.Expires,
		})
	}
	return fileInfos
}

func (e *MinioTemplate) PutObject(bucketName string, objectName string, data []byte) *file.FileObjectInfo {
	e.CreateBucket(bucketName)
	var minioClient = application.GetMinioClient()
	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewBuffer(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		logger.Error(err.Error())
	}
	return &file.FileObjectInfo{
		Name:         uploadInfo.Key,
		Size:         uploadInfo.Size,
		LastModified: uploadInfo.LastModified,
		Expires:      uploadInfo.Expiration,
	}
}
