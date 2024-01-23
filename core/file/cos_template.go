package file

import (
	"bytes"
	"context"
	"core/utils"
	"net/http"
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
	v, _, err := e.cosClient.Bucket.Get(context.Background(), opt)
	if err != nil {
		return nil, err
	}
	var fileInfos = make([]FileObjectInfo, 0)
	for _, c := range v.Contents {
		fileInfos = append(fileInfos, FileObjectInfo{
			Name:        c.Key,
			ContentType: c.ETag,
			Size:        c.Size,
		})
	}
	return fileInfos, nil
}

func (e *CosTemplate) PutObject(folderName string, objectName string, data []byte) (*FileObjectInfo, error) {
	if len(folderName) > 0 {
		objectName = folderName + utils.SLASH + objectName
	}
	_, err := e.cosClient.Object.Put(context.Background(), objectName, bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	return &FileObjectInfo{
		Name: objectName,
		Size: int64(len(data)),
	}, nil
}

func (e *CosTemplate) PresignedGetObject(folderName string, objectName string, headers map[string]string, expires time.Duration) (string, error) {
	if len(folderName) > 0 {
		objectName = folderName + utils.SLASH + objectName
	}
	credential := e.cosClient.GetCredential()
	presignedUrl, err := e.cosClient.Object.GetPresignedURL(context.Background(), http.MethodGet, objectName, credential.SecretID, credential.SecretKey, expires, nil)
	if err != nil {
		return "", err
	}
	if presignedUrl == nil {
		return "", nil
	}
	return presignedUrl.String(), nil
}
