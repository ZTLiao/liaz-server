package model

import (
	"core/types"
)

type AdminUser struct {
	AdminId      int64      `json:"admin_id,omitempty" xorm:"admin_id pk autoincr BIGINT"`
	Name         string     `json:"name,omitempty" xorm:"name"`
	Username     string     `json:"username,omitempty" xorm:"username"`
	Password     string     `json:"password,omitempty" xorm:"password"`
	Salt         string     `json:"salt,omitempty" xorm:"salt"`
	Phone        string     `json:"phone,omitempty" xorm:"phone"`
	Avatar       string     `json:"avatar,omitempty" xorm:"avatar"`
	Email        string     `json:"email,omitempty" xorm:"email"`
	Introduction string     `json:"introduction,omitempty" xorm:"introduction"`
	Status       int        `json:"status,omitempty" xorm:"status"`
	CreatedAt    types.Time `json:"createdAt,omitempty" xorm:"created_at timestamp created"`
	UpdatedAt    types.Time `json:"updatedAt,omitempty" xorm:"updated_at timestamp updated"`
}

func (e *AdminUser) TableName() string {
	return "admin_user"
}
