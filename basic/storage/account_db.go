package storage

import (
	"basic/enums"
	"basic/model"
	"core/constant"
	"core/errors"
	"core/types"
	"net/http"
	"time"

	"github.com/go-xorm/xorm"
)

type AccountDb struct {
	db *xorm.Engine
}

func NewAccountDb(db *xorm.Engine) *AccountDb {
	return &AccountDb{db}
}

func (e *AccountDb) IsExistForUsername(username string, password string) (bool, error) {
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

func (e *AccountDb) SignUpForUsername(username string, password string, flag int8) (int64, error) {
	var now = types.Time(time.Now())
	var account = new(model.Account)
	account.Username = username
	account.Password = password
	account.Flag = flag
	account.Status = enums.ACCOUNT_STATUS_OF_ENABLE
	account.CreatedAt = now
	account.UpdatedAt = now
	count, err := e.db.Where("username = ?", username).Count(account)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New(http.StatusForbidden, constant.USERNAME_EXISTS)
	}
	_, err = e.db.Insert(account)
	if err != nil {
		return 0, err
	}
	return account.UserId, nil
}

func (e *AccountDb) GetAccountById(userId int64) (*model.Account, error) {
	var account model.Account
	_, err := e.db.ID(userId).Get(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
