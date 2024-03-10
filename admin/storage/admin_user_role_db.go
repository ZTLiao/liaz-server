package storage

import (
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AdminUserRoleDb struct {
	db *xorm.Engine
}

func NewAdminUserRoleDb(db *xorm.Engine) *AdminUserRoleDb {
	return &AdminUserRoleDb{db}
}

func (e *AdminUserRoleDb) GetAdminUserRole(adminId int64) ([]model.AdminUserRole, error) {
	var adminUserRoles []model.AdminUserRole
	err := e.db.Where("admin_id = ?", adminId).Find(&adminUserRoles)
	if err != nil {
		return nil, err
	}
	return adminUserRoles, err
}

func (e *AdminUserRoleDb) DelAdminUserRole(adminId int64, roleId int64) error {
	if adminId == 0 && roleId == 0 {
		return nil
	}
	if adminId != 0 {
		_, err := e.db.Where("admin_id = ?", adminId).Delete(&model.AdminUserRole{})
		if err != nil {
			return err
		}
	}
	if roleId != 0 {
		_, err := e.db.Where("role_id = ?", roleId).Delete(&model.AdminUserRole{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AdminUserRoleDb) AddAdminUserRole(adminId int64, roleId int64) error {
	var adminUserRole = new(model.AdminUserRole)
	count, err := e.db.Where("admin_id = ? and role_id = ?", adminId, roleId).Count(&model.AdminUserRole{})
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	adminUserRole.AdminId = adminId
	adminUserRole.RoleId = roleId
	adminUserRole.CreatedAt = types.Time(time.Now())
	_, err = e.db.Insert(adminUserRole)
	if err != nil {
		return err
	}
	return nil
}
