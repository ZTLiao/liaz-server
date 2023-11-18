package controller

import (
	"admin/handler"
	"basic/infrastructure/persistence/repository"
	"basic/infrastructure/persistence/repository/memory"
	"basic/infrastructure/persistence/repository/store"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminSysConfController struct {
}

func (e *AdminSysConfController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var sysConfStore = store.NewSysConfStore(db)
	var sysConfMemory = memory.NewSysConfMemory(redis)
	var adminSysConfHandler = &handler.AdminSysConfHandler{
		SysConfApp: repository.NewSysConfRepo(sysConfStore, sysConfMemory),
	}
	iWebRoutes.GET("/sys/conf", adminSysConfHandler.GetAdminSysConf)
	iWebRoutes.POST("/sys/conf", adminSysConfHandler.SaveAdminSysConf)
	iWebRoutes.PUT("/sys/conf", adminSysConfHandler.UpdateAdminSysConf)
	iWebRoutes.DELETE("/sys/conf/:confId", adminSysConfHandler.DelAdminSysConf)
	iWebRoutes.GET("/sys/conf/:confKind", adminSysConfHandler.GetAdminSysConfByKind)
}
