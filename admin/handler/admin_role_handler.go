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

func (e *AdminRoleHandler) GetAdminRole(wc *web.WebContext) interface{} {
	adminRole, err := e.AdminRoleDb.GetAdminRole()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminRole)
}

func (e *AdminRoleHandler) SaveAdminRole(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminRole(wc)
	return response.Success()
}

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
