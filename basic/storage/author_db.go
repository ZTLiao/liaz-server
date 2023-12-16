package storage

import (
	"basic/model"

	"github.com/go-xorm/xorm"
)

type AuthorDb struct {
	db *xorm.Engine
}

func NewAuthorDb(db *xorm.Engine) *AuthorDb {
	return &AuthorDb{db}
}

func (e *AuthorDb) GetAuthorPage(startRow int, endRow int) ([]model.Author, int64, error) {
	var authors []model.Author
	err := e.db.OrderBy("seq_no asc").Limit(endRow, startRow).Find(&authors)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.Author{})
	if err != nil {
		return nil, 0, err
	}
	return authors, total, nil
}

func (e *AuthorDb) GetAuthorList() ([]model.Author, error) {
	var authors []model.Author
	err := e.db.OrderBy("seq_no asc").Find(&authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (e *AuthorDb) SaveOrUpdateAuthor(author *model.Author) error {
	authorId := author.AuthorId
	if authorId == 0 {
		_, err := e.db.Insert(author)
		if err != nil {
			return err
		}
	} else {
		_, err := e.db.ID(authorId).Update(author)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AuthorDb) DelAuthor(authorId int64) error {
	_, err := e.db.ID(authorId).Delete(&model.Author{})
	if err != nil {
		return err
	}
	return nil
}
