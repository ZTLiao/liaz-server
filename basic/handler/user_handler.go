package handler

import (
	"basic/device"
	"basic/resp"
	"basic/storage"
	"core/constant"
	"core/response"
	"core/web"
	"strconv"
)

type UserHandler struct {
	UserDb    *storage.UserDb
	AccountDb *storage.AccountDb
}

func (e *UserHandler) GetUser(wc *web.WebContext) interface{} {
	userIdStr := wc.Query("userId")
	if len(userIdStr) == 0 {
		return response.Success()
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	userResp, err := e.GetUserById(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(userResp)
}

func (e *UserHandler) GetUserById(userId int64) (*resp.UserResp, error) {
	user, err := e.UserDb.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	account, err := e.AccountDb.GetAccountById(userId)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}
	var userResp = &resp.UserResp{
		UserId:      userId,
		Username:    account.Username,
		Nickname:    user.Nickname,
		Phone:       user.Phone,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Description: user.Description,
		Gender:      user.Gender,
		Country:     user.Country,
		Province:    user.Province,
		City:        user.City,
	}
	return userResp, nil
}

func (e *UserHandler) UpdateUser(wc *web.WebContext) interface{} {
	userIdStr := wc.PostForm("userId")
	if len(userIdStr) == 0 {
		return response.Fail(constant.USERNAME_NOT_EXISTS)
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	nickname := wc.PostForm("nickname")
	if len(nickname) == 0 {
		return response.Fail(constant.NICKNAME_EMPTY)
	}
	phone := wc.PostForm("phone")
	email := wc.PostForm("email")
	genderStr := wc.PostForm("gender")
	var gender int64
	if len(genderStr) != 0 {
		gender, err = strconv.ParseInt(genderStr, 10, 8)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	description := wc.PostForm("description")
	err = e.UserDb.UpdateUser(userId, nickname, phone, email, int8(gender), description)
	if err != nil {
		wc.AbortWithError(err)
	}
	deviceInfo := device.GetDeviceInfo(wc)
	clientIp := deviceInfo.ClientIp
	if len(clientIp) != 0 {
		e.UserDb.UpdateLocation(userId, clientIp)
	}
	err = e.AccountDb.UpdatePhoneAndEmail(userId, phone, email)
	if err != nil {
		wc.AbortWithError(err)
	}
	userResp, err := e.GetUserById(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(userResp)
}
