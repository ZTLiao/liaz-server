package handler

import (
	"basic/storage"
	"core/constant"
	"core/response"
	"core/web"
)

type AccountHandler struct {
	AccountDb       *storage.AccountDb
	VerifyCodeCache *storage.VerifyCodeCache
}

func (e *AccountHandler) ResetPassword(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	verifyCode := wc.PostForm("verifyCode")
	newPassword := wc.PostForm("newPassword")
	wc.Info("username : %v, verifyCode : %v, newPassword : %v", username, verifyCode, newPassword)
	if len(username) == 0 {
		return response.Fail(constant.EMAIL_EMPTY)
	}
	if len(verifyCode) == 0 {
		return response.Fail(constant.VERIFY_CODE_EMPTY)
	}
	if len(newPassword) == 0 {
		return response.Fail(constant.NEW_PASSWORD_EMPTY)
	}
	account, err := e.AccountDb.GetAccountByUsername(username)
	if err != nil {
		wc.AbortWithError(err)
	}
	if account == nil {
		return response.Fail(constant.USERNAME_NOT_EXISTS)
	}
	userId := account.UserId
	err = e.AccountDb.ResetPassword(userId, newPassword)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.Success()
}
