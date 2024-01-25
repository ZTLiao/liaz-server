package model

import (
	"core/model"
	"core/types"
)

type AdminUser struct {
	AdminId      int64      `json:"adminId" xorm:"admin_id pk autoincr BIGINT"`
	Name         string     `json:"name" xorm:"name"`
	Username     string     `json:"username" xorm:"username"`
	Password     string     `json:"password" xorm:"password"`
	Salt         string     `json:"salt" xorm:"salt"`
	Phone        string     `json:"phone" xorm:"phone"`
	Avatar       string     `json:"avatar" xorm:"avatar"`
	Email        string     `json:"email" xorm:"email"`
	Introduction string     `json:"introduction" xorm:"introduction"`
	Status       int8       `json:"status" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt    types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &AdminUser{}

func (e *AdminUser) TableName() string {
	return "admin_user"
}
