package persistence

import (
	"basic/domain/entity"
	"basic/domain/repository"
	"core/constant"
	"core/logger"

	"github.com/go-xorm/xorm"
)

type FileItemRepo struct {
	db *xorm.Engine
}

var _ repository.FileItemRepository = &FileItemRepo{}

func NewFileItemRepository(db *xorm.Engine) *FileItemRepo {
	return &FileItemRepo{db}
}

func (e *FileItemRepo) SaveFileItem(fileItem *entity.FileItem) (*entity.FileItem, map[string]string) {
	dbErr := map[string]string{}
	_, err := e.db.Insert(fileItem)
	if err != nil {
		dbErr[constant.DB_ERROR] = err.Error()
		logger.Error(err.Error())
		return nil, dbErr
	}
	return fileItem, nil
}
func (e *FileItemRepo) GetFileItem(fileId int64) (*entity.FileItem, error) {
	var fileItem entity.FileItem
	_, err := e.db.ID(fileId).Get(&fileItem)
	if err != nil {
		return nil, err
	}
	return &fileItem, nil
}
func (e *FileItemRepo) DeleteFileItem(fileId int64) error {
	var fileItem entity.FileItem
	_, err := e.db.ID(fileId).Delete(&fileItem)
	if err != nil {
		return err
	}
	return nil
}
