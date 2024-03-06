package handler

import (
	basicStorage "basic/storage"
	"business/resp"
	businessStorage "business/storage"
	"core/constant"
	"core/response"
	"core/utils"
	"core/web"
	"strconv"
	"strings"
)

type DiscussHandler struct {
	DiscussDb         *businessStorage.DiscussDb
	DiscussResourceDb *businessStorage.DiscussResourceDb
	UserDb            *basicStorage.UserDb
}

func (e *DiscussHandler) Discuss(wc *web.WebContext) interface{} {
	parentIdStr := wc.PostForm("parentId")
	var parentId int64
	var err error
	if len(parentIdStr) != 0 {
		parentId, err = strconv.ParseInt(parentIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	userId := web.GetUserId(wc)
	objIdStr := wc.PostForm("objId")
	objId, err := strconv.ParseInt(objIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	objTypeStr := wc.PostForm("objType")
	objType, err := strconv.ParseInt(objTypeStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	content := wc.PostForm("content")
	discussId, err := e.DiscussDb.Save(parentId, userId, objId, int8(objType), content, constant.YES)
	if err != nil {
		wc.AbortWithError(err)
	}
	var resType int8
	resTypeStr := wc.PostForm("resType")
	if len(resTypeStr) != 0 {
		value, err := strconv.ParseInt(resTypeStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		resType = int8(value)
	}
	paths := wc.PostForm("paths")
	if len(paths) == 0 {
		return response.Success()
	}
	pathArray := strings.Split(paths, utils.COMMA)
	for _, path := range pathArray {
		err = e.DiscussResourceDb.Save(discussId, resType, path)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	return response.Success()
}

func (e *DiscussHandler) GetDiscussByObjId(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	objIdStr := wc.PostForm("objId")
	objId, err := strconv.ParseInt(objIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	objTypeStr := wc.PostForm("objType")
	objType, err := strconv.ParseInt(objTypeStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	wc.Info("pageNum : %v, pageSize : %v, objId : %v, objType : %v", pageNum, pageSize, objId, objType)
	discusses, err := e.DiscussDb.GetDiscussPage(objId, int8(objType), int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(discusses) == 0 {
		return response.Success()
	}
	var discussResps = make([]resp.DiscussResp, 0)
	for _, v := range discusses {
		discussResps = append(discussResps, resp.DiscussResp{
			DiscussId: v.DiscussId,
			UserId:    v.UserId,
		})
	}
	return response.ReturnOK(discussResps)
}
