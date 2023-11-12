package model

import "core/types"

type AdminUserRole struct {
	AdminId   int64      `json:"adminId" xorm:"admin_id"`
	RoleId    int64      `json:"roleId" xorm:"role_id"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at"`
}

func (e *AdminUserRole) TableName() string {
	return "admin_user_role"
}
