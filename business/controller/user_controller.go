package controller

import (
	"basic/handler"
	"basic/storage"
	"core/system"
	"core/web"
)

type UserController struct {
}

var _ web.IWebController = &UserController{}

func (e *UserController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var userHandler = &handler.UserHandler{
		UserDb:    storage.NewUserDb(db),
		AccountDb: storage.NewAccountDb(db),
	}
	iWebRoutes.GET("/user/get", userHandler.GetUser)
	iWebRoutes.POST("/user/update", userHandler.UpdateUser)
}
