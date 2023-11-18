package repository

import "basic/infrastructure/persistence/entity"

type FileItemRepoInterface interface {
	SaveFileItem(*entity.FileItem) (*entity.FileItem, error)
	GetFileItem(int64) (*entity.FileItem, error)
	DeleteFileItem(int64) error
}
