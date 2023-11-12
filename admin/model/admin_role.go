package model

import "core/types"

type AdminRole struct {
	RoleId    int64      `json:"roleId" xorm:"role_id pk autoincr BIGINT"`
	Name      string     `json:"name" xorm:"name"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *AdminRole) TableName() string {
	return "admin_role"
}
