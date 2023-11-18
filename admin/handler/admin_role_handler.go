package handler

import (
	"admin/model"
	"admin/storage"
	"core/response"
	"core/web"
	"fmt"
	"strconv"
)

type AdminRoleHandler struct {
	AdminRoleDb     *storage.AdminRoleDb
	AdminRoleMenuDb *storage.AdminRoleMenuDb
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
	adminRole, err := e.AdminRoleDb.GetAdminRole()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminRole)
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
	if err := wc.ShouldBindJSON(&params); err != nil {
		wc.AbortWithError(err)
	}
	roleIdStr := fmt.Sprint(params["roleId"])
	name := fmt.Sprint(params["name"])
	var adminRole = new(model.AdminRole)
	if len(roleIdStr) > 0 {
		roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		adminRole.RoleId = roleId
	}
	adminRole.Name = name
	e.AdminRoleDb.SaveOrUpdateAdminRole(adminRole)
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
	roleIdStr := wc.Param("roleId")
	if len(roleIdStr) > 0 {
		roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AdminRoleDb.DelAdminRole(roleId)
		e.AdminRoleMenuDb.DelAdminRoleMenu(roleId, 0)
	}
	return response.Success()
}
