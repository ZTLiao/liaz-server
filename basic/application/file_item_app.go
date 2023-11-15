package application

import (
	"basic/domain/entity"
	"basic/domain/repository"
)

type fileItemApp struct {
	fileItemRepo repository.FileItemRepository
}

var _ FileItemAppInterface = &fileItemApp{}

type FileItemAppInterface interface {
	SaveFileItem(*entity.FileItem) (*entity.FileItem, map[string]string)
	GetFileItem(int64) (*entity.FileItem, error)
	DeleteFileItem(int64) error
}

func (e *fileItemApp) SaveFileItem(fileItem *entity.FileItem) (*entity.FileItem, map[string]string) {
	return e.fileItemRepo.SaveFileItem(fileItem)
}
func (e *fileItemApp) GetFileItem(fileId int64) (*entity.FileItem, error) {
	return e.fileItemRepo.GetFileItem(fileId)
}
func (e *fileItemApp) DeleteFileItem(fileId int64) error {
	return e.fileItemRepo.DeleteFileItem(fileId)
}
