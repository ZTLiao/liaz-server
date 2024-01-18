package handler

import (
	"basic/device"
	"basic/enums"
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
	if e.Email == nil {
		return response.Fail(constant.SERVER_ERROR)
	}
	subject := e.Email.Subject[constant.VERIFY_CODE_TEMPLATE]
	template := e.Email.Template[constant.VERIFY_CODE_TEMPLATE]
	randomNumber := strconv.FormatInt(int64(utils.RandomForSix()), 10)
	body := fmt.Sprintf(template, randomNumber)
	wc.Info("username : %v, host : %v, port : %v, nickname : %v, emailTo : %v, subject : %v, body : %v", e.Email.Username, e.Email.Host, e.Email.Port, e.Email.Nickname, emailTo, subject, body)
	err := utils.SendMail(e.Email.Username, e.Email.Password, e.Email.Host, e.Email.Port, e.Email.Nickname, emailTo, subject, body)
	var resCode = http.StatusOK
	var resMsg string
	if err != nil {
		resCode = http.StatusInternalServerError
		resMsg = err.Error()
	}
	wc.Info("resCode : %v, resMsg : %v", resCode, resMsg)
	e.VerifyCodeRecordDb.InsertVerifyCodeRecord(emailTo, enums.SEND_TYPE_FOR_EMAIL, randomNumber, strconv.FormatInt(int64(resCode), 10), resMsg, device.GetDeviceInfo(wc))
	if resCode == http.StatusOK {
		e.VerifyCodeCache.Set(emailTo, randomNumber)
	}
	return response.Success()
}

func (e *VerifyCodeHandler) CheckVerifyCode(wc *web.WebContext) interface{} {
	username := wc.PostForm("username")
	verifyCode := wc.PostForm("verifyCode")
	cacheCode, err := e.VerifyCodeCache.Get(username)
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(cacheCode) == 0 {
		return response.ReturnOK(false)
	}
	return response.ReturnOK(verifyCode == cacheCode)
}
