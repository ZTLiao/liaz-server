package handler

import (
	"basic/interfaces"
	"core/web"
)

type AdminUploadHandler struct {
	FileItemHandler *interfaces.FileItemHandler
}

// @Summary 上传文件
// @title Swagger API
// @Tags 文件管理
// @description 上传文件接口
// @BasePath /admin/upload/:bucketName
// @Produce json
// @Param roleId query int64 true "桶名称"
// @Param file formData string true "文件"
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/upload/:bucketName [post]
func (e *AdminUploadHandler) Upload(wc *web.WebContext) interface{} {
	return e.FileItemHandler.UploadFile(wc)
}
