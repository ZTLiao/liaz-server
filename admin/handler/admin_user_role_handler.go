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
	AdminUserRoleDb *storage.AdminUserRoleDb
	AdminRoleDb     *storage.AdminRoleDb
}

func (e *AdminUserRoleHandler) GetAdminUserRole(wc *web.WebContext) interface{} {
	adminIdStr := wc.Param("adminId")
	if len(adminIdStr) == 0 {
		return response.Success()
	}
	adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminUserRoles, err := e.AdminUserRoleDb.GetAdminUserRole(adminId)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminRoles, err := e.AdminRoleDb.GetAdminRole()
	if err != nil {
		wc.AbortWithError(err)
	}
	var roles = make([]resp.AdminRoleResp, 0)
	for i := 0; i < len(adminRoles); i++ {
		role := adminRoles[i]
		var checked bool
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

func (e *AdminUserRoleHandler) SaveAdminUserRole(wc *web.WebContext) interface{} {
	adminIdStr := wc.PostForm("adminId")
	roleIds := wc.PostForm("roleIds")
	wc.Info("adminId : %s, roleIds : %s", adminIdStr, roleIds)
	if len(adminIdStr) == 0 {
		return response.Success()
	}
	adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	e.AdminUserRoleDb.DelAdminUserRole(adminId, 0)
	if len(roleIds) > 0 {
		roleIdArray := strings.Split(roleIds, utils.COMMA)
		for i := 0; i < len(roleIdArray); i++ {
			roleIdStr := roleIdArray[i]
			roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			e.AdminUserRoleDb.AddAdminUserRole(adminId, roleId)
		}
	}
	return response.Success()
}
