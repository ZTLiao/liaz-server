package controller

import (
	"basic/handler"
	"basic/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AccountController struct {
}

var _ web.IWebController = &AccountController{}

func (e *AccountController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var accountHandler = handler.AccountHandler{
		AccountDb:       storage.NewAccountDb(db),
		VerifyCodeCache: storage.NewVerifyCodeCache(redis),
	}
	iWebRoutes.POST("/account/reset/password", accountHandler.ResetPassword)
}
