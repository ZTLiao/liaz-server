package storage

import "github.com/go-xorm/xorm"

type AccountDb struct {
	db *xorm.Engine
}

func NewAccountDb(db *xorm.Engine) *AccountDb {
	return &AccountDb{db}
}
