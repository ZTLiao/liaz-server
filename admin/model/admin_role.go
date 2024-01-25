package model

import (
	"core/model"
	"core/types"
)

type AdminRole struct {
	RoleId    int64      `json:"roleId" xorm:"role_id pk autoincr BIGINT"`
	Name      string     `json:"name" xorm:"name"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &AdminRole{}

func (e *AdminRole) TableName() string {
	return "admin_role"
}
