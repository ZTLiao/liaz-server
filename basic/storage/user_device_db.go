package storage

import "github.com/go-xorm/xorm"

type UserDeviceDb struct {
	db *xorm.Engine
}

func NewUserDeviceDb(db *xorm.Engine) *UserDeviceDb {
	return &UserDeviceDb{db}
}
