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

type AdminRoleMenuHandler struct {
	AdminRoleMenuDb storage.AdminRoleMenuDb
	AdminMenuDb     storage.AdminMenuDb
}

// @Summary 获取角色菜单
// @title Swagger API
// @Tags 角色管理
// @description 获取角色菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/role/menu/:roleId
// @Param roleId query int64 true "角色ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role/menu/:roleId [get]
func (e *AdminRoleMenuHandler) GetAdminRoleMenu(wc *web.WebContext) interface{} {
	var roleIdStr = wc.Context.Param("roleId")
	if len(roleIdStr) == 0 {
		return response.Success()
	}
	roleId, _ := strconv.ParseInt(roleIdStr, 10, 64)
	var adminRoleMenus = e.AdminRoleMenuDb.GetAdminRoleMenu(roleId)
	var adminMenus = e.AdminMenuDb.GetAdminMenuList(wc.Background())
	var menus = make([]resp.AdminMenuResp, 0)
	var childMap = make(map[int64][]resp.AdminMenuResp, 0)
	for i := 0; i < len(adminMenus); i++ {
		var adminMenu = adminMenus[i]
		var menuId = adminMenu.MenuId
		var parentId = adminMenu.ParentId
		var checked bool = false
		for _, v := range adminRoleMenus {
			if v.MenuId == menuId {
				checked = true
				break
			}
		}
		var menu = resp.AdminMenuResp{
			MenuId:   menuId,
			MenuName: adminMenu.Name,
			ParentId: parentId,
			Checked:  checked,
		}
		if parentId == 0 {
			menus = append(menus, menu)
		} else {
			var childs = childMap[parentId]
			if len(childs) == 0 {
				childs = make([]resp.AdminMenuResp, 0)
			}
			childMap[parentId] = append(childs, menu)
		}
	}
	for i := 0; i < len(menus); i++ {
		var menu = &menus[i]
		var menuId = menu.MenuId
		var childs []resp.AdminMenuResp = childMap[menuId]
		if len(childs) == 0 {
			continue
		}
		menu.Childs = make([]resp.AdminMenuResp, 0)
		menu.Childs = append(menu.Childs, childs...)
	}
	return response.ReturnOK(menus)
}

// @Summary 授权
// @title Swagger API
// @Tags 角色管理
// @description 授权接口
// @BasePath /admin/role/menu
// @Produce json
// @Param roleId formData string true "角色ID"
// @Param menuIds formData string true "多个菜单ID"
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/role/menu [post]
func (e *AdminRoleMenuHandler) SaveAdminRoleMenu(wc *web.WebContext) interface{} {
	var roleIdStr = wc.Context.PostForm("roleId")
	var menuIds = wc.Context.PostForm("menuIds")
	wc.Info("roleId : %s, menuIds : %s", roleIdStr, menuIds)
	if len(roleIdStr) == 0 {
		return response.Success()
	}
	roleId, _ := strconv.ParseInt(roleIdStr, 10, 64)
	e.AdminRoleMenuDb.DelAdminRoleMenu(roleId, 0)
	if len(menuIds) > 0 {
		var menuIdArray = strings.Split(menuIds, utils.COMMA)
		for i := 0; i < len(menuIdArray); i++ {
			var menuIdStr = menuIdArray[i]
			menuId, _ := strconv.ParseInt(menuIdStr, 10, 64)
			e.AdminRoleMenuDb.AddAdminRoleMenu(roleId, menuId)
		}
	}
	return response.Success()
}
