package storage

import (
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminUserRoleDb struct {
}

func (e *AdminUserRoleDb) GetAdminUserRole(adminId int64) []model.AdminUserRole {
	var engine = application.GetApp().GetXormEngine()
	var adminUserRoles []model.AdminUserRole
	err := engine.Where("admin_id = ?", adminId).Find(&adminUserRoles)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminUserRoles
}

func (e *AdminUserRoleDb) DelAdminUserRole(adminId int64, roleId int64) {
	if adminId == 0 && roleId == 0 {
		return
	}
	var engine = application.GetApp().GetXormEngine()
	if adminId != 0 {
		engine.Where("admin_id = ?", adminId).Delete(&model.AdminUserRole{})
	}
	if roleId != 0 {
		engine.Where("role_id = ?", roleId).Delete(&model.AdminUserRole{})
	}
}

func (e *AdminUserRoleDb) AddAdminUserRole(adminId int64, roleId int64) {
	var engine = application.GetApp().GetXormEngine()
	var adminUserRole = new(model.AdminUserRole)
	count, err := engine.Where("admin_id = ? and role_id = ?", adminId, roleId).Count(adminUserRole)
	if err != nil {
		logger.Error(err.Error())
	}
	if count > 0 {
		return
	}
	adminUserRole.AdminId = adminId
	adminUserRole.RoleId = roleId
	engine.Insert(adminUserRole)
}
