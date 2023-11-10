package storage

import (
	"admin/enums"
	"admin/model"
	"core/application"
	"core/logger"
)

type AdminUserStorage struct {
}

// 获取登录用户信息
func (e *AdminUserStorage) GetLoginUser(username string, password string) *model.AdminUser {
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
