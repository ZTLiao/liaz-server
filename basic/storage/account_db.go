package storage

import (
	"basic/enums"
	"basic/model"

	"github.com/go-xorm/xorm"
)

type AccountDb struct {
	db *xorm.Engine
}

func NewAccountDb(db *xorm.Engine) *AccountDb {
	return &AccountDb{db}
}

func (e *AccountDb) IsExist(username string, password string) (bool, error) {
	count, err := e.db.Where("(username = ? or phone = ? or email = ?) and password = ? and status = ?", username, username, username, password, enums.ACCOUNT_STATUS_OF_ENABLE).Count(&model.Account{})
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return count == 1, nil
}

func (e *AccountDb) SignIn(username string, password string) (*model.Account, error) {
	var accounts []model.Account
	err := e.db.Where("(username = ? or phone = ? or email = ?) and password = ? and status = ?", username, username, username, password, enums.ACCOUNT_STATUS_OF_ENABLE).Find(&accounts)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, nil
	}
	return &accounts[0], nil
}
