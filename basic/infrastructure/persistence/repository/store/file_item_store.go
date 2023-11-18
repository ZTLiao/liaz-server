package store

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"

	"github.com/go-xorm/xorm"
)

type FileItemStore struct {
	db *xorm.Engine
}

var _ repository.FileItemRepoInterface = &FileItemStore{}

func NewFileItemStore(db *xorm.Engine) *FileItemStore {
	return &FileItemStore{db}
}

func (e *FileItemStore) SaveFileItem(fileItem *entity.FileItem) (*entity.FileItem, error) {
	_, err := e.db.Insert(fileItem)
	if err != nil {
		return nil, err
	}
	_, err = e.db.Where("file_id = ?", fileItem.FileId).Get(fileItem)
	if err != nil {
		return nil, err
	}
	return fileItem, nil
}

func (e *FileItemStore) GetFileItem(fileId int64) (*entity.FileItem, error) {
	var fileItem entity.FileItem
	_, err := e.db.ID(fileId).Get(&fileItem)
	if err != nil {
		return nil, err
	}
	return &fileItem, nil
}

func (e *FileItemStore) DeleteFileItem(fileId int64) error {
	var fileItem entity.FileItem
	_, err := e.db.ID(fileId).Delete(&fileItem)
	if err != nil {
		return err
	}
	return nil
}
