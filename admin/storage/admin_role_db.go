package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminRoleDb struct {
}

func (e *AdminRoleDb) GetAdminRole() []model.AdminRole {
	var engine = application.GetApp().GetXormEngine()
	var adminRoles []model.AdminRole
	err := engine.OrderBy("created_at asc").Find(&adminRoles)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminRoles
}

func (e *AdminRoleDb) SaveOrUpdateAdminRole(adminRole *model.AdminRole) {
	var engine = application.GetApp().GetXormEngine()
	var roleId = adminRole.RoleId
	var name = adminRole.Name
	if roleId == 0 {
		count, err := engine.Where("name = ?", name).Count(adminRole)
		if err != nil {
			logger.Error(err.Error())
		}
		if count == 0 {
			engine.Insert(adminRole)
		}
	} else {
		engine.ID(roleId).Update(adminRole)
	}
}

func (e *AdminRoleDb) DelAdminRole(roleId int64) {
	var engine = application.GetApp().GetXormEngine()
	engine.ID(roleId).Delete(&model.AdminRole{})
}
