package entity

import "core/types"

type FileItem struct {
	FileId     int64      `json:"fileId" xorm:"file_id pk autoincr BIGINT"`
	BucketName string     `json:"bucketName" xorm:"bucket_name"`
	ObjectName string     `json:"objectName" xorm:"object_name"`
	Size       int64      `json:"size" xorm:"size"`
	Path       string     `json:"path" xorm:"path"`
	UnqiueId   string     `json:"unqiueId" xorm:"unqiue_id"`
	Suffix     string     `json:"suffix" xorm:"suffix"`
	FileType   string     `json:"fileType" xorm:"file_type"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt  types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *FileItem) TableName() string {
	return "file_item"
}
