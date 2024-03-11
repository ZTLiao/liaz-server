package model

import (
	"core/model"
	"core/types"
)

type FileItem struct {
	FileId     int64      `json:"fileId" xorm:"file_id pk autoincr BIGINT"`
	BucketName string     `json:"bucketName" xorm:"bucket_name"`
	ObjectName string     `json:"objectName" xorm:"object_name"`
	Size       int64      `json:"size" xorm:"size"`
	Path       string     `json:"path" xorm:"path"`
	UniqueId   string     `json:"uniqueId" xorm:"unique_id"`
	Suffix     string     `json:"suffix" xorm:"suffix"`
	FileType   string     `json:"fileType" xorm:"file_type"`
	Status     int8       `json:"status" xorm:"status"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt  types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &FileItem{}

func (e *FileItem) TableName() string {
	return "file_item"
}
