package file

import (
	"bytes"
	"core/utils"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssTemplate struct {
	ossClient *oss.Client
}

var _ FileTemplate = &OssTemplate{}

func NewOssTemplate(ossClient *oss.Client) *OssTemplate {
	return &OssTemplate{ossClient}
}

func (e *OssTemplate) CreateBucket(bucketName string) error {
	ok, err := e.ossClient.IsBucketExist(bucketName)
	if err != nil {
		return err
	}
	if !ok {
		err := e.ossClient.CreateBucket(bucketName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *OssTemplate) ListObjects(bucketName string) ([]FileObjectInfo, error) {
	bucket, err := e.ossClient.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	var marker = utils.EMPTY
	var fileInfos = make([]FileObjectInfo, 0)
	for {
		result, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return nil, err
		}
		for _, object := range result.Objects {
			fileInfos = append(fileInfos, FileObjectInfo{
				Name:         object.Key,
				ContentType:  object.Type,
				Size:         object.Size,
				LastModified: object.LastModified,
			})
		}
		if result.IsTruncated {
			marker = result.NextMarker
		} else {
			break
		}
	}
	return fileInfos, nil
}

func (e *OssTemplate) PutObject(bucketName string, objectName string, data []byte) (*FileObjectInfo, error) {
	e.CreateBucket(bucketName)
	bucket, err := e.ossClient.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	err = bucket.PutObject(objectName, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return &FileObjectInfo{
		Name: objectName,
		Size: int64(len(data)),
	}, nil
}

func (e *OssTemplate) PresignedGetObject(bucketName string, objectName string, headers map[string]string, expires time.Duration) (string, error) {
	bucket, err := e.ossClient.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	url, err := bucket.SignURL(objectName, oss.HTTPGet, int64(expires.Seconds()))
	if err != nil {
		return "", nil
	}
	return url, nil
}
