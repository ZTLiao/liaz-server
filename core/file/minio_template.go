package file

import (
	"bytes"
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioTemplate struct {
	minioClient *minio.Client
}

var _ FileTemplate = &MinioTemplate{}

func NewMinioTemplate(minioClient *minio.Client) *MinioTemplate {
	return &MinioTemplate{minioClient}
}

func (e *MinioTemplate) CreateBucket(bucketName string) error {
	ok, err := e.minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}
	if !ok {
		err := e.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *MinioTemplate) ListObjects(bucketName string) ([]FileObjectInfo, error) {
	objectInfos := e.minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
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
	return fileInfos, nil
}

func (e *MinioTemplate) PutObject(bucketName string, objectName string, data []byte) (*FileObjectInfo, error) {
	e.CreateBucket(bucketName)
	uploadInfo, err := e.minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewBuffer(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &FileObjectInfo{
		Name:         uploadInfo.Key,
		Size:         uploadInfo.Size,
		LastModified: uploadInfo.LastModified,
		Expires:      uploadInfo.Expiration,
	}, nil
}

func (e *MinioTemplate) PresignedGetObject(bucketName string, objectName string, headers map[string]string, expires time.Duration) (string, error) {
	var reqParams = make(url.Values)
	for k, v := range headers {
		reqParams.Set(k, v)
	}
	presignedURL, err := e.minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expires, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.RequestURI(), err
}
