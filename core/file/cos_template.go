package file

import (
	"context"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosTemplate struct {
	cosClient *cos.Client
}

var _ FileTemplate = &CosTemplate{}

func NewCosTemplate(cosClient *cos.Client) *CosTemplate {
	return &CosTemplate{cosClient}
}

func (e *CosTemplate) CreateBucket(bucketName string) error {
	_, err := e.cosClient.Bucket.Put(context.Background(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *CosTemplate) ListObjects(folderName string) ([]FileObjectInfo, error) {
	opt := &cos.BucketGetOptions{
		Prefix:  folderName,
		MaxKeys: 3,
	}
	e.cosClient.Bucket.Get(context.Background(), opt)
	return nil, nil
}

func (e *CosTemplate) PutObject(folderName string, objectName string, data []byte) (*FileObjectInfo, error) {
	return nil, nil
}

func (e *CosTemplate) PresignedGetObject(folderName string, objectName string, headers map[string]string, expires time.Duration) (string, error) {
	return "", nil
}
