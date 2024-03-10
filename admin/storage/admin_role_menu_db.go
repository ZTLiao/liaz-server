package storage

import (
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AdminRoleMenuDb struct {
	db *xorm.Engine
}

func NewAdminRoleMenuDb(db *xorm.Engine) *AdminRoleMenuDb {
	return &AdminRoleMenuDb{db}
}

func (e *AdminRoleMenuDb) AddAdminRoleMenu(roleId int64, menuId int64) error {
	var adminRoleMenu = new(model.AdminRoleMenu)
	count, err := e.db.Where("role_id = ? and menu_id = ?", roleId, menuId).Count(&model.AdminRoleMenu{})
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	adminRoleMenu.RoleId = roleId
	adminRoleMenu.MenuId = menuId
	adminRoleMenu.CreatedAt = types.Time(time.Now())
	_, err = e.db.Insert(adminRoleMenu)
	if err != nil {
		return err
	}
	return nil
}

func (e *AdminRoleMenuDb) DelAdminRoleMenu(roleId int64, menuId int64) error {
	if roleId == 0 && menuId == 0 {
		return nil
	}
	if roleId != 0 {
		_, err := e.db.Where("role_id = ?", roleId).Delete(&model.AdminRoleMenu{})
		if err != nil {
			return err
		}
	}
	if menuId != 0 {
		_, err := e.db.Where("menu_id = ?", menuId).Delete(&model.AdminRoleMenu{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AdminRoleMenuDb) GetAdminRoleMenu(roleId int64) ([]model.AdminRoleMenu, error) {
	var adminRoleMenus []model.AdminRoleMenu
	err := e.db.Where("role_id = ?", roleId).Find(&adminRoleMenus)
	if err != nil {
		return nil, err
	}
	return adminRoleMenus, nil
}
