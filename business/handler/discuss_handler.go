package handler

import (
	basicStorage "basic/storage"
	"business/enums"
	"business/resp"
	businessStorage "business/storage"
	"core/constant"
	"core/event"
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
	if objType == enums.ASSET_TYPE_FOR_COMIC {
		event.Bus.Publish(constant.COMIC_DISCUSS_RANK_TOPIC, objId)
	} else if objType == enums.ASSET_TYPE_FOR_NOVEL {
		event.Bus.Publish(constant.NOVEL_DISCUSS_RANK_TOPIC, objId)
	}
	return response.Success()
}

func (e *DiscussHandler) GetDiscussPage(wc *web.WebContext) interface{} {
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
		discussId := v.DiscussId
		userId := v.UserId
		user, err := e.UserDb.GetUserById(userId)
		if err != nil {
			wc.AbortWithError(err)
		}
		discussResources, err := e.DiscussResourceDb.GetDiscussResourceByDiscussId(discussId)
		if err != nil {
			wc.AbortWithError(err)
		}
		var paths = make([]string, 0)
		if len(discussResources) != 0 {
			for _, v := range discussResources {
				paths = append(paths, v.Path)
			}
		}
		parentId := v.ParentId
		var parents = make([]resp.DiscussResp, 0)
		err = e.BuildParentDiscuss(parentId, &parents)
		if err != nil {
			wc.AbortWithError(err)
		}
		discussResps = append(discussResps, resp.DiscussResp{
			DiscussId: discussId,
			UserId:    userId,
			CreatedAt: v.CreatedAt,
			Content:   v.Content,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Gender:    user.Gender,
			Paths:     paths,
		})
	}
	return response.ReturnOK(discussResps)
}

func (e *DiscussHandler) BuildParentDiscuss(discussId int64, parents *[]resp.DiscussResp) error {
	if discussId == 0 {
		return nil
	}
	discuss, err := e.DiscussDb.GetDiscussById(discussId)
	if err != nil {
		return err
	}
	userId := discuss.UserId
	user, err := e.UserDb.GetUserById(userId)
	if err != nil {
		return err
	}
	discussResources, err := e.DiscussResourceDb.GetDiscussResourceByDiscussId(discussId)
	if err != nil {
		return err
	}
	var paths = make([]string, 0)
	if len(discussResources) != 0 {
		for _, v := range discussResources {
			paths = append(paths, v.Path)
		}
	}
	parentId := discuss.ParentId
	err = e.BuildParentDiscuss(parentId, parents)
	if err != nil {
		return err
	}
	var discussResp = resp.DiscussResp{
		DiscussId: discussId,
		UserId:    userId,
		CreatedAt: discuss.CreatedAt,
		Content:   discuss.Content,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Paths:     paths,
	}
	*parents = append(*parents, discussResp)
	return nil
}
