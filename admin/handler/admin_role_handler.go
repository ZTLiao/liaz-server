package handler

import (
	"admin/model"
	"admin/storage"
	"core/errors"
	"core/response"
	"core/web"
	"fmt"
	"strconv"
)

type AdminRoleHandler struct {
	AdminRoleDb     storage.AdminRoleDb
	AdminRoleMenuDb storage.AdminRoleMenuDb
}

// @Summary 获取系统所有角色
// @title Swagger API
// @Tags 角色管理
// @description 获取系统所有角色接口
// @Security ApiKeyAuth
// @BasePath /admin/role
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role [get]
func (e *AdminRoleHandler) GetAdminRole(wc *web.WebContext) interface{} {
	return response.ReturnOK(e.AdminRoleDb.GetAdminRole(wc.Background()))
}

// @Summary 保存角色
// @title Swagger API
// @Tags 角色管理
// @description 保存角色接口
// @Security ApiKeyAuth
// @BasePath /admin/role
// @Param adminRole body model.AdminRole true "角色"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role [post]
func (e *AdminRoleHandler) SaveAdminRole(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminRole(wc)
	return response.Success()
}

// @Summary 修改角色
// @title Swagger API
// @Tags 角色管理
// @description 修改角色接口
// @Security ApiKeyAuth
// @BasePath /admin/role
// @Param adminRole body model.AdminRole true "角色"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role [put]
func (e *AdminRoleHandler) UpdateAdminRole(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminRole(wc)
	return response.Success()
}

func (e *AdminRoleHandler) saveOrUpdateAdminRole(wc *web.WebContext) {
	var params map[string]any
	if err := wc.Context.ShouldBindJSON(&params); err != nil {
		wc.Context.Error(&errors.ApiError{
			Message: err.Error(),
		})
		return
	}
	var roleId = fmt.Sprint(params["roleId"])
	var name = fmt.Sprint(params["name"])
	var adminRole = new(model.AdminRole)
	if len(roleId) > 0 {
		adminRole.RoleId, _ = strconv.ParseInt(roleId, 10, 64)
	}
	adminRole.Name = name
	e.AdminRoleDb.SaveOrUpdateAdminRole(wc.Background(), adminRole)
}

// @Summary 删除角色
// @title Swagger API
// @Tags 角色管理
// @description 删除角色接口
// @Security ApiKeyAuth
// @BasePath /admin/role/:roleId
// @Param roleId query int64 true "角色ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role/:roleId [delete]
func (e *AdminRoleHandler) DelAdminRole(wc *web.WebContext) interface{} {
	var roleIdStr = wc.Context.Param("roleId")
	if len(roleIdStr) > 0 {
		roleId, _ := strconv.ParseInt(roleIdStr, 10, 64)
		e.AdminRoleDb.DelAdminRole(wc.Background(), roleId)
		e.AdminRoleMenuDb.DelAdminRoleMenu(wc.Background(), roleId, 0)
	}
	return response.Success()
}
