package handler

import (
	"admin/resp"
	"admin/storage"
	"core/response"
	"core/utils"
	"core/web"
	"strconv"
	"strings"
)

type AdminUserRoleHandler struct {
	AdminUserRoleDb storage.AdminUserRoleDb
	AdminRoleDb     storage.AdminRoleDb
}

// @Summary 获取用户角色
// @title Swagger API
// @Tags 角色管理
// @description 获取角色菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/user/role/:adminId
// @Param adminId query int64 true "用户ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user/role/:adminId [get]
func (e *AdminUserRoleHandler) GetAdminUserRole(wc *web.WebContext) interface{} {
	var adminIdStr = wc.Context.Param("adminId")
	if len(adminIdStr) == 0 {
		return response.Success()
	}
	adminId, _ := strconv.ParseInt(adminIdStr, 10, 64)
	var adminUserRoles = e.AdminUserRoleDb.GetAdminUserRole(wc.Background(), adminId)
	var adminRoles = e.AdminRoleDb.GetAdminRole(wc.Background())
	var roles = make([]resp.AdminRoleResp, 0)
	for i := 0; i < len(adminRoles); i++ {
		var role = adminRoles[i]
		var checked bool = false
		for _, v := range adminUserRoles {
			if role.RoleId == v.RoleId {
				checked = true
				break
			}
		}
		roles = append(roles, resp.AdminRoleResp{
			RoleId:   role.RoleId,
			RoleName: role.Name,
			Checked:  checked,
		})
	}
	return response.ReturnOK(roles)
}

// @Summary 分配角色
// @title Swagger API
// @Tags 用户管理
// @description 分配角色接口
// @BasePath /admin/user/role
// @Produce json
// @Param adminId formData string true "用户ID"
// @Param roleIds formData string true "多个角色ID"
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/user/role [post]
func (e *AdminUserRoleHandler) SaveAdminUserRole(wc *web.WebContext) interface{} {
	var adminIdStr = wc.Context.PostForm("adminId")
	var roleIds = wc.Context.PostForm("roleIds")
	wc.Info("adminId : %s, roleIds : %s", adminIdStr, roleIds)
	if len(adminIdStr) == 0 {
		return response.Success()
	}
	adminId, _ := strconv.ParseInt(adminIdStr, 10, 64)
	e.AdminUserRoleDb.DelAdminUserRole(wc.Background(), adminId, 0)
	if len(roleIds) > 0 {
		var roleIdArray = strings.Split(roleIds, utils.COMMA)
		for i := 0; i < len(roleIdArray); i++ {
			var roleIdStr = roleIdArray[i]
			roleId, _ := strconv.ParseInt(roleIdStr, 10, 64)
			e.AdminUserRoleDb.AddAdminUserRole(wc.Background(), adminId, roleId)
		}
	}
	return response.Success()
}
