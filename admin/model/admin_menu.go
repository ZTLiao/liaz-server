package model

import (
	"core/model"
	"core/types"
)

type AdminMenu struct {
	MenuId      int64      `json:"menuId" xorm:"menu_id pk autoincr BIGINT"`
	ParentId    int64      `json:"parentId" xorm:"parent_id"`
	Name        string     `json:"name" xorm:"name"`
	Path        string     `json:"path" xorm:"path"`
	Icon        string     `json:"icon" xorm:"icon"`
	Status      int8       `json:"status" xorm:"status"`
	ShowOrder   int        `json:"showOrder" xorm:"show_order"`
	Description string     `json:"description" xorm:"description"`
	CreatedAt   types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt   types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &AdminMenu{}

func (e *AdminMenu) TableName() string {
	return "admin_menu"
}
