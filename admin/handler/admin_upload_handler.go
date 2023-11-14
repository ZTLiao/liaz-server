package handler

import (
	"core/file/minio"
	"core/response"
	"core/web"
	"io"
)

type AdminUploadHandler struct {
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
	var bucketName = wc.Context.Param("bucketName")
	if len(bucketName) == 0 {
		return response.Success()
	}
	file, header, err := wc.Context.Request.FormFile("file")
	if err != nil {
		return response.Fail(err.Error())
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return response.Fail(err.Error())
	}
	var fileInfo = new(minio.MinioTemplate).PutObject(bucketName, header.Filename, data)
	return response.ReturnOK(&map[string]string{
		"path": fileInfo.Name,
	})
}
