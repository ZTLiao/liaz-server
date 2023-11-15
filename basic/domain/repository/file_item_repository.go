package repository

import "basic/domain/entity"

type FileItemRepository interface {
	SaveFileItem(*entity.FileItem) (*entity.FileItem, map[string]string)
	GetFileItem(int64) (*entity.FileItem, error)
	DeleteFileItem(int64) error
}
