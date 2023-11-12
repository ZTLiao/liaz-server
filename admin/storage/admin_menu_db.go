package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminMenuDb struct {
}

func (e *AdminMenuDb) GetAdminMemu(adminId int64) []model.AdminMenu {
	var engine = application.GetApp().GetXormEngine()
	var adminMenus []model.AdminMenu
	err := engine.SQL(
		`
		select 
			am.menu_id,
			am.parent_id,
			am.name,
			am.path,
			am.icon,
			am.show_order,
			am.description,
			am.status,
			am.created_at,
			am.updated_at
		from admin_user_role as aur 
		inner join admin_role_menu as arm on arm.role_id = aur.role_id
		inner join admin_menu as am on am.menu_id = arm.menu_id
		where 
			aur.admin_id = ?
		group by am.menu_id
		order by am.parent_id, am.show_order
		`, adminId).Find(&adminMenus)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminMenus
}

func (e *AdminMenuDb) GetAdminMenuList() []model.AdminMenu {
	var engine = application.GetApp().GetXormEngine()
	var adminMenus []model.AdminMenu
	err := engine.OrderBy("created_at asc").Find(&adminMenus)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminMenus
}

func (e *AdminMenuDb) SaveOrUpdateAdminMenu(adminMenu *model.AdminMenu) {
	var engine = application.GetApp().GetXormEngine()
	var menuId = adminMenu.MenuId
	var name = adminMenu.Name
	var path = adminMenu.Path
	if menuId == 0 {
		count, err := engine.Where("name = ? and path = ?", name, path).Count(adminMenu)
		if err != nil {
			logger.Error(err.Error())
		}
		if count == 0 {
			engine.Insert(adminMenu)
		}
	} else {
		engine.ID(menuId).Update(adminMenu)
	}
	_, err := engine.Where("name = ? and path = ?", name, path).Get(adminMenu)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (e *AdminMenuDb) DelAdminMenu(menuId int64) {
	var engine = application.GetApp().GetXormEngine()
	engine.ID(menuId).Delete(&model.AdminMenu{})
}
