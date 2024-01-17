package handler

import (
	"basic/resp"
	"basic/storage"
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
	user, err := e.UserDb.GetUserById(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if user == nil {
		return response.Success()
	}
	account, err := e.AccountDb.GetAccountById(userId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if account == nil {
		return response.Success()
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
	return response.ReturnOK(userResp)
}
