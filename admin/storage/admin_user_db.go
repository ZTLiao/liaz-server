package storage

import (
	"admin/enums"
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminUserDb struct {
}

// 获取登录用户信息
func (e *AdminUserDb) GetLoginUser(username string, password string) *model.AdminUser {
	var engine = application.GetApp().GetXormEngine()
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

func (e *AdminUserDb) GetAdminUserList() []model.AdminUser {
	var engine = application.GetApp().GetXormEngine()
	var adminUsers []model.AdminUser
	err := engine.OrderBy("created_at desc").Find(&adminUsers)
	if err != nil {
		logger.Error(err.Error())
	}
	return adminUsers
}

func (e *AdminUserDb) SaveOrUpdateAdminUser(adminUser *model.AdminUser) {
	var engine = application.GetApp().GetXormEngine()
	var adminId = adminUser.AdminId
	if adminId == 0 {
		engine.Insert(adminUser)
	} else {
		engine.ID(adminId).Update(adminUser)
	}
}

func (e *AdminUserDb) DelAdminUser(adminId int64) {
	var engine = application.GetApp().GetXormEngine()
	engine.ID(adminId).Delete(&model.AdminUser{})
}

func (e *AdminUserDb) ThawAdminUser(adminId int64) {
	var engine = application.GetApp().GetXormEngine()
	engine.ID(adminId).Update(&model.AdminUser{
		Status: 1,
	})
}
