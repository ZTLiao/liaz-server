package handler

import (
	"basic/storage"
	"core/constant"
	"core/response"
	"core/web"
)

type AccountHandler struct {
	AccountDb *storage.AccountDb
}

func (e *AccountHandler) ResetPassword(wc *web.WebContext) interface{} {
	email := wc.PostForm("email")
	verifyCode := wc.PostForm("verifyCode")
	newPassword := wc.PostForm("newPassword")
	wc.Info("email : %v, verifyCode : %v, newPassword : %v", email, verifyCode, newPassword)
	if len(email) == 0 {
		return response.Fail(constant.EMAIL_EMPTY)
	}
	if len(verifyCode) == 0 {
		return response.Fail(constant.VERIFY_CODE_EMPTY)
	}
	if len(newPassword) == 0 {
		return response.Fail(constant.NEW_PASSWORD_EMPTY)
	}
	return response.Success()
}
