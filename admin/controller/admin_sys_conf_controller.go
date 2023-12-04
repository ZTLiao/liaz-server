package controller

import (
	admin "admin/handler"
	basic "basic/handler"
	"basic/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminSysConfController struct {
}

func (e *AdminSysConfController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var sysConfDb = storage.NewSysConfDb(db)
	var sysConfCache = storage.NewSysConfCache(redis)
	var adminSysConfHandler = &admin.AdminSysConfHandler{
		SysConfHandler: basic.NewSysConfHandler(sysConfDb, sysConfCache),
	}
	iWebRoutes.GET("/sys/conf", adminSysConfHandler.GetAdminSysConf)
	iWebRoutes.POST("/sys/conf", adminSysConfHandler.SaveAdminSysConf)
	iWebRoutes.PUT("/sys/conf", adminSysConfHandler.UpdateAdminSysConf)
	iWebRoutes.DELETE("/sys/conf/:confId", adminSysConfHandler.DelAdminSysConf)
	iWebRoutes.GET("/sys/conf/:confKind", adminSysConfHandler.GetAdminSysConfByKind)
}
