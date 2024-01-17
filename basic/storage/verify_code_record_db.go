package storage

import "github.com/go-xorm/xorm"

type VerifyCodeRecordDb struct {
	db *xorm.Engine
}

func NewVerifyCodeRecordDb(db *xorm.Engine) *VerifyCodeRecordDb {
	return &VerifyCodeRecordDb{db}
}
