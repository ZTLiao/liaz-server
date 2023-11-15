package storage

import (
	"admin/model"
	"context"
	"core/application"
	"core/logger"
)

type AdminUserRoleDb struct {
}

func (e *AdminUserRoleDb) GetAdminUserRole(ctx context.Context, adminId int64) []model.AdminUserRole {
	var engine = application.GetXormEngine()
	var adminUserRoles []model.AdminUserRole
	err := engine.Where("admin_id = ?", adminId).Find(&adminUserRoles)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminUserRoles
}

func (e *AdminUserRoleDb) DelAdminUserRole(ctx context.Context, adminId int64, roleId int64) {
	if adminId == 0 && roleId == 0 {
		return
	}
	var engine = application.GetXormEngine()
	if adminId != 0 {
		engine.Where("admin_id = ?", adminId).Delete(&model.AdminUserRole{})
	}
	if roleId != 0 {
		engine.Where("role_id = ?", roleId).Delete(&model.AdminUserRole{})
	}
}

func (e *AdminUserRoleDb) AddAdminUserRole(ctx context.Context, adminId int64, roleId int64) {
	var engine = application.GetXormEngine()
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
