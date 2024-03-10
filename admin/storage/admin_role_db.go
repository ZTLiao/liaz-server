package storage

import (
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AdminRoleDb struct {
	db *xorm.Engine
}

func NewAdminRoleDb(db *xorm.Engine) *AdminRoleDb {
	return &AdminRoleDb{db}
}

func (e *AdminRoleDb) GetAdminRole() ([]model.AdminRole, error) {
	var adminRoles []model.AdminRole
	err := e.db.OrderBy("created_at asc").Find(&adminRoles)
	if err != nil {
		return nil, err
	}
	return adminRoles, nil
}

func (e *AdminRoleDb) SaveOrUpdateAdminRole(adminRole *model.AdminRole) error {
	var now = types.Time(time.Now())
	roleId := adminRole.RoleId
	name := adminRole.Name
	if roleId == 0 {
		count, err := e.db.Where("name = ?", name).Count(&model.AdminRole{})
		if err != nil {
			return err
		}
		if count == 0 {
			adminRole.CreatedAt = now
			_, err := e.db.Insert(adminRole)
			if err != nil {
				return err
			}
		}
	} else {
		adminRole.UpdatedAt = now
		_, err := e.db.ID(roleId).Update(adminRole)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AdminRoleDb) DelAdminRole(roleId int64) error {
	_, err := e.db.ID(roleId).Delete(&model.AdminRole{})
	if err != nil {
		return err
	}
	return nil
}
