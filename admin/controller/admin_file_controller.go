package controller

import (
	adminHandler "admin/handler"
	basicHandler "basic/handler"
	"basic/storage"
	"core/constant"
	"core/file"
	"core/redis"
	"core/system"
	"core/web"
)

type AdminFileController struct {
}

var _ web.IWebController = &AdminFileController{}

func (e *AdminFileController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var fileItemDb = storage.NewFileItemDb(db)
	var fileTemplate = file.NewFileTemplate(constant.COS)
	var sysConfDb = storage.NewSysConfDb(db)
	var sysConfCache = storage.NewSysConfCache(redis)
	var adminFileHandler = adminHandler.AdminFileHandler{
		FileItemHandler: basicHandler.NewFileItemHandler(fileItemDb, fileTemplate),
		SysConfHandler:  basicHandler.NewSysConfHandler(sysConfDb, sysConfCache),
	}
	iWebRoutes.POST("/upload/:bucketName", adminFileHandler.Upload)
	iWebRoutes.POST("/upload/batch/:bucketName", adminFileHandler.UploadBatch)
	iWebRoutes.GET("/file/:bucketName/:objectName", adminFileHandler.PresignedGetObject)
}
