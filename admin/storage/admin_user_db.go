package storage

import (
	"admin/enums"
	"admin/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AdminUserDb struct {
	db *xorm.Engine
}

func NewAdminUserDb(db *xorm.Engine) *AdminUserDb {
	return &AdminUserDb{db}
}

// 获取登录用户信息
func (e *AdminUserDb) GetLoginUser(username string, password string) (*model.AdminUser, error) {
	var adminUsers []model.AdminUser
	err := e.db.Where("(username = ? or phone = ? or email = ?) and password = ? and status = ?", username, username, username, password, enums.USER_STATUS_OF_ENABLE).Find(&adminUsers)
	if err != nil {
		return nil, err
	}
	if len(adminUsers) == 0 {
		return nil, nil
	}
	return &adminUsers[0], nil
}

func (e *AdminUserDb) GetAdminUserList() ([]model.AdminUser, error) {
	var adminUsers []model.AdminUser
	err := e.db.OrderBy("created_at desc").Find(&adminUsers)
	if err != nil {
		return nil, err
	}
	return adminUsers, nil
}

func (e *AdminUserDb) SaveOrUpdateAdminUser(adminUser *model.AdminUser) error {
	var now = types.Time(time.Now())
	adminId := adminUser.AdminId
	if adminId == 0 {
		adminUser.CreatedAt = now
		_, err := e.db.Insert(adminUser)
		if err != nil {
			return err
		}
	} else {
		adminUser.UpdatedAt = now
		_, err := e.db.ID(adminId).Update(adminUser)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AdminUserDb) DelAdminUser(adminId int64) error {
	_, err := e.db.ID(adminId).Delete(&model.AdminUser{})
	if err != nil {
		return err
	}
	return nil
}

func (e *AdminUserDb) ThawAdminUser(adminId int64) error {
	_, err := e.db.ID(adminId).Update(&model.AdminUser{
		Status: 1,
	})
	if err != nil {
		return err
	}
	return nil
}
