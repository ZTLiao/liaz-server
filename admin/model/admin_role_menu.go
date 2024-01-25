package model

import (
	"core/model"
	"core/types"
)

type AdminRoleMenu struct {
	RoleId    int64      `json:"roleId" xorm:"role_id"`
	MenuId    int64      `json:"menuId" xorm:"menu_id"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &AdminRoleMenu{}

func (e *AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}
