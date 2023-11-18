package handler

import (
	"admin/model"
	"admin/storage"
	"core/constant"
	"core/logger"
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

// @Summary 获取系统所有菜单
// @title Swagger API
// @Tags 菜单管理
// @description 获取系统所有菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/menu/list
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/menu/list [get]
func (e *AdminMenuHandler) GetAdminMenuList(wc *web.WebContext) interface{} {
	adminMenus, err := e.AdminMenuDb.GetAdminMenuList()
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(adminMenus)
}

// @Summary 获取当前用户菜单
// @title Swagger API
// @Tags 首页管理
// @description 获取当前用户菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/menu
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/menu [get]
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

// @Summary 保存菜单
// @title Swagger API
// @Tags 菜单管理
// @description 保存菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/menu
// @Param adminMenu body model.AdminMenu true "菜单"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/menu [post]
func (e *AdminMenuHandler) SaveAdminMenu(wc *web.WebContext) interface{} {
	e.saveOrUpdateAdminMenu(wc)
	return response.Success()
}

// @Summary 修改菜单
// @title Swagger API
// @Tags 菜单管理
// @description 修改菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/menu
// @Param adminMenu body model.AdminMenu true "菜单"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/menu [put]
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

// @Summary 删除菜单
// @title Swagger API
// @Tags 菜单管理
// @description 删除菜单接口
// @Security ApiKeyAuth
// @BasePath /admin/menu/:menuId
// @Param menuId query int64 true "菜单ID"
// @Produce json
// @Success 200 {object} response.Response "{"code":200,"data":{},"message":"OK"}"
// @Router /admin/menu/:menuId [delete]
func (e *AdminMenuHandler) DelAdminMenu(wc *web.WebContext) interface{} {
	menuIdStr := wc.Param("menuId")
	if len(menuIdStr) > 0 {
		menuId, err := strconv.ParseInt(menuIdStr, 10, 64)
		if err != nil {
			logger.Error(err.Error())
		}
		e.AdminMenuDb.DelAdminMenu(menuId)
		e.AdminRoleMenuDb.DelAdminRoleMenu(0, menuId)
	}
	return response.Success()
}
