package model

import (
	"core/model"
	"core/types"
)

type AdminUserRole struct {
	AdminId   int64      `json:"adminId" xorm:"admin_id"`
	RoleId    int64      `json:"roleId" xorm:"role_id"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &AdminUserRole{}

func (e *AdminUserRole) TableName() string {
	return "admin_user_role"
}
