package application

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"
)

type fileItemApp struct {
	fileItemRepo repository.FileItemRepoInterface
}

var _ FileItemAppInterface = &fileItemApp{}

type FileItemAppInterface interface {
	SaveFileItem(*entity.FileItem) (*entity.FileItem, error)
	GetFileItem(int64) (*entity.FileItem, error)
	DeleteFileItem(int64) error
}

func (e *fileItemApp) SaveFileItem(fileItem *entity.FileItem) (*entity.FileItem, error) {
	return e.fileItemRepo.SaveFileItem(fileItem)
}

func (e *fileItemApp) GetFileItem(fileId int64) (*entity.FileItem, error) {
	return e.fileItemRepo.GetFileItem(fileId)
}

func (e *fileItemApp) DeleteFileItem(fileId int64) error {
	return e.fileItemRepo.DeleteFileItem(fileId)
}
