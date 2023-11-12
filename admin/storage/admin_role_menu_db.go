package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminRoleMenuDb struct {
}

func (e *AdminRoleMenuDb) AddAdminRoleMenu(roleId int64, menuId int64) {
	var engine = application.GetApp().GetXormEngine()
	var adminRoleMenu = new(model.AdminRoleMenu)
	count, err := engine.Where("role_id = ? and menu_id = ?", roleId, menuId).Count(adminRoleMenu)
	if err != nil {
		logger.Error(err.Error())
	}
	if count > 0 {
		return
	}
	adminRoleMenu.RoleId = roleId
	adminRoleMenu.MenuId = menuId
	engine.Insert(adminRoleMenu)
}

func (e *AdminRoleMenuDb) DelAdminRoleMenu(roleId int64, menuId int64) {
	if roleId == 0 && menuId == 0 {
		return
	}
	var engine = application.GetApp().GetXormEngine()
	if roleId != 0 {
		engine.Where("role_id = ?", roleId).Delete(&model.AdminRoleMenu{})
	}
	if menuId != 0 {
		engine.Where("menu_id = ?", menuId).Delete(&model.AdminRoleMenu{})
	}
}

func (e *AdminRoleMenuDb) GetAdminRoleMenu(roleId int64) []model.AdminRoleMenu {
	var engine = application.GetApp().GetXormEngine()
	var adminRoleMenus []model.AdminRoleMenu
	engine.Where("role_id = ?", roleId).Find(&adminRoleMenus)
	return adminRoleMenus
}
