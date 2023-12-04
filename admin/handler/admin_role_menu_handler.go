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
	AdminRoleMenuDb *storage.AdminRoleMenuDb
	AdminMenuDb     *storage.AdminMenuDb
}

func (e *AdminRoleMenuHandler) GetAdminRoleMenu(wc *web.WebContext) interface{} {
	roleIdStr := wc.Param("roleId")
	if len(roleIdStr) == 0 {
		return response.Success()
	}
	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminRoleMenus, err := e.AdminRoleMenuDb.GetAdminRoleMenu(roleId)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminMenus, err := e.AdminMenuDb.GetAdminMenuList()
	if err != nil {
		wc.AbortWithError(err)
	}
	var menus = make([]resp.AdminMenuResp, 0)
	var childMap = make(map[int64][]resp.AdminMenuResp, 0)
	for i := 0; i < len(adminMenus); i++ {
		adminMenu := adminMenus[i]
		menuId := adminMenu.MenuId
		parentId := adminMenu.ParentId
		var checked bool
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
			childs := childMap[parentId]
			if len(childs) == 0 {
				childs = make([]resp.AdminMenuResp, 0)
			}
			childMap[parentId] = append(childs, menu)
		}
	}
	for i := 0; i < len(menus); i++ {
		menu := &menus[i]
		menuId := menu.MenuId
		childs := childMap[menuId]
		if len(childs) == 0 {
			continue
		}
		menu.Childs = make([]resp.AdminMenuResp, 0)
		menu.Childs = append(menu.Childs, childs...)
	}
	return response.ReturnOK(menus)
}

func (e *AdminRoleMenuHandler) SaveAdminRoleMenu(wc *web.WebContext) interface{} {
	roleIdStr := wc.PostForm("roleId")
	menuIds := wc.PostForm("menuIds")
	wc.Info("roleId : %s, menuIds : %s", roleIdStr, menuIds)
	if len(roleIdStr) == 0 {
		return response.Success()
	}
	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	e.AdminRoleMenuDb.DelAdminRoleMenu(roleId, 0)
	if len(menuIds) > 0 {
		menuIdArray := strings.Split(menuIds, utils.COMMA)
		for i := 0; i < len(menuIdArray); i++ {
			menuIdStr := menuIdArray[i]
			menuId, err := strconv.ParseInt(menuIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			e.AdminRoleMenuDb.AddAdminRoleMenu(roleId, menuId)
		}
	}
	return response.Success()
}
