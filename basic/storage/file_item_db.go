package storage

import (
	"basic/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type FileItemDb struct {
	db *xorm.Engine
}

func NewFileItemDb(db *xorm.Engine) *FileItemDb {
	return &FileItemDb{db}
}

func (e *FileItemDb) SaveFileItem(fileItem *model.FileItem) (*model.FileItem, error) {
	var now = types.Time(time.Now())
	fileItem.CreatedAt = now
	fileItem.UpdatedAt = now
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
	if fileItem.FileId == 0 {
		return nil, nil
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

func (e *FileItemDb) GetFileTypeByPath(path string) (string, error) {
	var fileItem model.FileItem
	_, err := e.db.Where("path = ?", path).Get(&fileItem)
	if err != nil {
		return "", err
	}
	if fileItem.FileId == 0 {
		return "", nil
	}
	return fileItem.FileType, nil
}

func (e *FileItemDb) GetFileItemByPath(path string) (*model.FileItem, error) {
	var fileItem model.FileItem
	_, err := e.db.Where("path = ?", path).Get(&fileItem)
	if err != nil {
		return nil, err
	}
	if fileItem.FileId == 0 {
		return nil, nil
	}
	return &fileItem, nil
}
