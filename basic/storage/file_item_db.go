package storage

import (
	"basic/model"

	"github.com/go-xorm/xorm"
)

type FileItemDb struct {
	db *xorm.Engine
}

func NewFileItemDb(db *xorm.Engine) *FileItemDb {
	return &FileItemDb{db}
}

func (e *FileItemDb) SaveFileItem(fileItem *model.FileItem) (*model.FileItem, error) {
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

func (e *FileItemDb) GetFileItem(fileId int64) (*model.FileItem, error) {
	var fileItem model.FileItem
	_, err := e.db.ID(fileId).Get(&fileItem)
	if err != nil {
		return nil, err
	}
	return &fileItem, nil
}

func (e *FileItemDb) DeleteFileItem(fileId int64) error {
	var fileItem model.FileItem
	_, err := e.db.ID(fileId).Delete(&fileItem)
	if err != nil {
		return err
	}
	return nil
}
