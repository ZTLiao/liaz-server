package handler

import (
	"admin/model"
	"admin/storage"
	"core/constant"
	"core/errors"
	"core/response"
	"core/web"
	"fmt"
	"net/http"
	"strconv"
)

type AdminMenuHandler struct {
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
	return response.ReturnOK(new(storage.AdminMenuDb).GetAdminMenuList())
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
	var accessToken = wc.Context.Request.Header.Get(constant.AUTHORIZATION)
	var adminUser = new(storage.AdminUserCache).Get(accessToken)
	if adminUser == nil {
		return response.ReturnError(http.StatusForbidden, constant.ILLEGAL_REQUEST)
	}
	return response.ReturnOK(new(storage.AdminMenuDb).GetAdminMemu(adminUser.AdminId))
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
	if err := wc.Context.ShouldBindJSON(&params); err != nil {
		wc.Context.Error(&errors.ApiError{
			Message: err.Error(),
		})
		return
	}
	var menuId = fmt.Sprint(params["menuId"])
	var parentId = fmt.Sprint(params["parentId"])
	var name = fmt.Sprint(params["name"])
	var path = fmt.Sprint(params["path"])
	var icon = fmt.Sprint(params["icon"])
	var statusStr = fmt.Sprint(params["status"])
	var showOrderStr = fmt.Sprint(params["showOrder"])
	var description = fmt.Sprint(params["description"])
	var adminMenu = new(model.AdminMenu)
	if len(menuId) > 0 {
		adminMenu.MenuId, _ = strconv.ParseInt(menuId, 10, 64)
	}
	if len(parentId) > 0 {
		adminMenu.ParentId, _ = strconv.ParseInt(parentId, 10, 64)
	}
	adminMenu.Name = name
	adminMenu.Path = path
	adminMenu.Icon = icon
	status, _ := strconv.ParseInt(statusStr, 10, 64)
	adminMenu.Status = int8(status)
	showOrder, _ := strconv.ParseInt(showOrderStr, 10, 32)
	adminMenu.ShowOrder = int(showOrder)
	adminMenu.Description = description
	new(storage.AdminMenuDb).SaveOrUpdateAdminMenu(adminMenu)
	wc.Info("menuId : %d", adminMenu.MenuId)
	new(storage.AdminRoleMenuDb).AddAdminRoleMenu(constant.SUPER_ADMIN, adminMenu.MenuId)
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
	var menuId = wc.Context.Param("menuId")
	if len(menuId) > 0 {
		val, _ := strconv.ParseInt(menuId, 10, 64)
		new(storage.AdminMenuDb).DelAdminMenu(val)
		new(storage.AdminRoleMenuDb).DelAdminRoleMenu(0, val)
	}
	return response.Success()
}
