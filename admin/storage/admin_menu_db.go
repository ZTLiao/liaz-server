package storage

import (
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AdminMenuDb struct {
	db *xorm.Engine
}

func NewAdminMenuDb(db *xorm.Engine) *AdminMenuDb {
	return &AdminMenuDb{db}
}

func (e *AdminMenuDb) GetAdminMemu(adminId int64) ([]model.AdminMenu, error) {
	var adminMenus []model.AdminMenu
	err := e.db.SQL(
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
		left join admin_role_menu as arm on arm.role_id = aur.role_id
		left join admin_menu as am on am.menu_id = arm.menu_id
		where 
			aur.admin_id = ?
			and am.menu_id is not null
		group by am.menu_id
		order by am.parent_id, am.show_order
		`, adminId).Find(&adminMenus)
	if err != nil {
		return nil, err
	}
	return adminMenus, nil
}

func (e *AdminMenuDb) GetAdminMenuList() ([]model.AdminMenu, error) {
	var adminMenus []model.AdminMenu
	err := e.db.OrderBy("created_at asc").Find(&adminMenus)
	if err != nil {
		return nil, err
	}
	return adminMenus, nil
}

func (e *AdminMenuDb) SaveOrUpdateAdminMenu(adminMenu *model.AdminMenu) error {
	var now = types.Time(time.Now())
	menuId := adminMenu.MenuId
	name := adminMenu.Name
	path := adminMenu.Path
	if menuId == 0 {
		count, err := e.db.Where("name = ? and path = ?", name, path).Count(&model.AdminMenu{})
		if err != nil {
			return err
		}
		if count == 0 {
			adminMenu.CreatedAt = now
			_, err := e.db.Insert(adminMenu)
			if err != nil {
				return err
			}
		}
	} else {
		adminMenu.UpdatedAt = now
		_, err := e.db.ID(menuId).Update(adminMenu)
		if err != nil {
			return err
		}
	}
	_, err := e.db.Where("name = ? and path = ?", name, path).Get(adminMenu)
	if err != nil {
		return err
	}
	return nil
}

func (e *AdminMenuDb) DelAdminMenu(menuId int64) error {
	_, err := e.db.ID(menuId).Delete(&model.AdminMenu{})
	if err != nil {
		return err
	}
	return nil
}
