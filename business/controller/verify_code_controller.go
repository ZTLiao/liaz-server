package controller

import (
	"basic/handler"
	"basic/storage"
	"core/config"
	"core/redis"
	"core/system"
	"core/web"
)

type VerifyCodeController struct {
}

var _ web.IWebController = &VerifyCodeController{}

func (e *VerifyCodeController) Router(iWebRoutes web.IWebRoutes) {
	email := config.SystemConfig.Email
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var verifyCodeHandler = handler.VerifyCodeHandler{
		Email:              email,
		VerifyCodeRecordDb: storage.NewVerifyCodeRecordDb(db),
		VerifyCodeCache:    storage.NewVerifyCodeCache(redis),
	}
	iWebRoutes.POST("/verify/code/email", verifyCodeHandler.SendVerifyCodeForEmail)
}
