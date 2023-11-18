package repository

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"
	"basic/infrastructure/persistence/repository/store"
)

type FileItemRepo struct {
	fileItemStore *store.FileItemStore
}

var _ repository.FileItemRepoInterface = &FileItemRepo{}

func NewFileItemRepo(fileItemStore *store.FileItemStore) *FileItemRepo {
	return &FileItemRepo{fileItemStore}
}

func (e *FileItemRepo) SaveFileItem(fileItem *entity.FileItem) (*entity.FileItem, error) {
	return e.fileItemStore.SaveFileItem(fileItem)
}

func (e *FileItemRepo) GetFileItem(fileId int64) (*entity.FileItem, error) {
	return e.fileItemStore.GetFileItem(fileId)
}

func (e *FileItemRepo) DeleteFileItem(fileId int64) error {
	return e.fileItemStore.DeleteFileItem(fileId)
}
