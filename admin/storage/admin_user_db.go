package storage

import (
	"admin/enums"
	"admin/model"
	"context"
	"core/application"
	"core/logger"
)

type AdminUserDb struct {
}

// 获取登录用户信息
func (e *AdminUserDb) GetLoginUser(ctx context.Context, username string, password string) *model.AdminUser {
	var engine = application.GetXormEngine()
	var adminUsers []model.AdminUser
	err := engine.Where("(username = ? or phone = ? or email = ?) and password = ? and status = ?", username, username, username, password, enums.USER_STATUS_OF_ENABLE).Find(&adminUsers)
	if err != nil {
		logger.Error(err.Error())
	}
	if len(adminUsers) == 0 {
		return nil
	}
	return &adminUsers[0]
}

func (e *AdminUserDb) GetAdminUserList(ctx context.Context) []model.AdminUser {
	var engine = application.GetXormEngine()
	var adminUsers []model.AdminUser
	err := engine.OrderBy("created_at desc").Find(&adminUsers)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminUsers
}

func (e *AdminUserDb) SaveOrUpdateAdminUser(ctx context.Context, adminUser *model.AdminUser) {
	var engine = application.GetXormEngine()
	var adminId = adminUser.AdminId
	if adminId == 0 {
		engine.Insert(adminUser)
	} else {
		engine.ID(adminId).Update(adminUser)
	}
}

func (e *AdminUserDb) DelAdminUser(ctx context.Context, adminId int64) {
	var engine = application.GetXormEngine()
	engine.ID(adminId).Delete(&model.AdminUser{})
}

func (e *AdminUserDb) ThawAdminUser(ctx context.Context, adminId int64) {
	var engine = application.GetXormEngine()
	engine.ID(adminId).Update(&model.AdminUser{
		Status: 1,
	})
}
