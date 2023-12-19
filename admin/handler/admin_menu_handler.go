package handler

import (
	"admin/model"
	"admin/storage"
	"core/constant"
	"core/response"
	"core/web"
	"fmt"
	"net/http"
	"strconv"
)

type AdminMenuHandler struct {
	AdminMenuDb     *storage.AdminMenuDb
	AdminRoleMenuDb *storage.AdminRoleMenuDb
	AdminUserCache  *storage.AdminUserCache
}

func (e *AdminMenuHandler) GetAdminMenuList(wc *web.WebContext) interface{} {
	adminMenus, err := e.AdminMenuDb.GetAdminMenuList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminMenus)
}

func (e *AdminMenuHandler) GetAdminMenu(wc *web.WebContext) interface{} {
	accessToken := wc.GetHeader(constant.AUTHORIZATION)
	adminUser, err := e.AdminUserCache.Get(accessToken)
	if err != nil {
		wc.AbortWithError(err)
	}
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	adminMenu, err := e.AdminMenuDb.GetAdminMemu(adminUser.AdminId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminMenu)
}

func (e *AdminMenuHandler) SaveAdminMenu(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminMenu(wc)
	return response.Success()
}

func (e *AdminMenuHandler) UpdateAdminMenu(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminMenu(wc)
	return response.Success()
}

func (e *AdminMenuHandler) saveOrUpdateAdminMenu(wc *web.WebContext) {
	var params map[string]any
	if err := wc.ShouldBindJSON(&params); err != nil {
		wc.AbortWithError(err)
	}
	menuIdStr := fmt.Sprint(params["menuId"])
	parentIdStr := fmt.Sprint(params["parentId"])
	name := fmt.Sprint(params["name"])
	path := fmt.Sprint(params["path"])
	icon := fmt.Sprint(params["icon"])
	statusStr := fmt.Sprint(params["status"])
	showOrderStr := fmt.Sprint(params["showOrder"])
	description := fmt.Sprint(params["description"])
	var adminMenu = new(model.AdminMenu)
	if len(menuIdStr) > 0 {
		menuId, err := strconv.ParseInt(menuIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		adminMenu.MenuId = menuId
	}
	if len(parentIdStr) > 0 {
		parentId, err := strconv.ParseInt(parentIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		adminMenu.ParentId = parentId
	}
	adminMenu.Name = name
	adminMenu.Path = path
	adminMenu.Icon = icon
	status, err := strconv.ParseInt(statusStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminMenu.Status = int8(status)
	showOrder, err := strconv.ParseInt(showOrderStr, 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	adminMenu.ShowOrder = int(showOrder)
	adminMenu.Description = description
	e.AdminMenuDb.SaveOrUpdateAdminMenu(adminMenu)
	wc.Info("menuId : %d", adminMenu.MenuId)
	e.AdminRoleMenuDb.AddAdminRoleMenu(constant.SUPER_ADMIN, adminMenu.MenuId)
}

func (e *AdminMenuHandler) DelAdminMenu(wc *web.WebContext) interface{} {
	menuIdStr := wc.Param("menuId")
	if len(menuIdStr) > 0 {
		menuId, err := strconv.ParseInt(menuIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		e.AdminMenuDb.DelAdminMenu(menuId)
		e.AdminRoleMenuDb.DelAdminRoleMenu(0, menuId)
	}
	return response.Success()
}
