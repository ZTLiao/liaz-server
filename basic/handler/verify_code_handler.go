package handler

import (
	"basic/storage"
	"core/config"
	"core/constant"
	"core/response"
	"core/utils"
	"core/web"
	"fmt"
	"net/http"
	"strconv"
)

type VerifyCodeHandler struct {
	Email              *config.Email
	VerifyCodeRecordDb *storage.VerifyCodeRecordDb
	VerifyCodeCache    *storage.VerifyCodeCache
}

func (e *VerifyCodeHandler) SendVerifyCodeForEmail(wc *web.WebContext) interface{} {
	emailTo := wc.PostForm("email")
	email := e.Email
	if email == nil {
		return response.Fail(constant.SERVER_ERROR)
	}
	subject := email.Subject[constant.VERIFY_CODE_TEMPLATE]
	template := email.Template[constant.VERIFY_CODE_TEMPLATE]
	randomNumber := strconv.FormatInt(int64(utils.RandomForSix()), 10)
	body := fmt.Sprintf(template, randomNumber)
	err := utils.SendMail(email.Username, email.Password, email.Host, email.Port, email.Nickname, emailTo, subject, body)
	var resCode = http.StatusOK
	var resMsg string
	if err != nil {
		resCode = http.StatusInternalServerError
		resMsg = err.Error()
	}
	wc.Info("resCode : %v, resMsg : %v", resCode, resMsg)
	if resCode == http.StatusOK {
		e.VerifyCodeCache.Set(emailTo, randomNumber)
	}
	return response.Success()
}
