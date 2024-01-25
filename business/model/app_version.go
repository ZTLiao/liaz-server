package model

import (
	"core/model"
	"core/types"
)

type AppVersion struct {
	VersionId    int64      `json:"versionId" xorm:"version_id pk autoincr BIGINT"`
	Os           string     `json:"os" xorm:"os"`
	Version      string     `json:"version" xorm:"version"`
	Channel      string     `json:"channel" xorm:"channel"`
	Description  string     `json:"description" xorm:"description"`
	DownloadLink string     `json:"downloadLink" xorm:"download_link"`
	FileMd5      string     `json:"fileMd5" xorm:"file_md5"`
	Status       int8       `json:"status" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt    types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &AppVersion{}

func (e *AppVersion) TableName() string {
	return "app_version"
}
