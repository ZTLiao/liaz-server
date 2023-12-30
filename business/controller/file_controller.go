package controller

import (
	basicHandler "basic/handler"
	"basic/storage"
	businessHandler "business/handler"
	"core/file"
	"core/redis"
	"core/system"
	"core/web"
)

type FileController struct {
}

var _ web.IWebController = &FileController{}

func (e *FileController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	minio := system.GetMinioClient()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var fileItemDb = storage.NewFileItemDb(db)
	var fileTemplate = file.NewFileTemplate(minio)
	var sysConfDb = storage.NewSysConfDb(db)
	var sysConfCache = storage.NewSysConfCache(redis)
	var fileHandler = &businessHandler.FileHandler{
		FileItemHandler: basicHandler.NewFileItemHandler(fileItemDb, fileTemplate),
		SysConfHandler:  basicHandler.NewSysConfHandler(sysConfDb, sysConfCache),
	}
	iWebRoutes.POST("/file/upload", fileHandler.Upload)
	iWebRoutes.GET("/file/:bucketName/:objectName", fileHandler.PresignedGetObject)
}
