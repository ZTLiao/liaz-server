package storage

import "github.com/go-xorm/xorm"

type UserDb struct {
	db *xorm.Engine
}

func NewUserDb(db *xorm.Engine) *UserDb {
	return &UserDb{db}
}
