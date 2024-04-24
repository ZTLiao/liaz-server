package handler

import (
	"admin/resp"
	"business/model"
	"business/storage"
	"core/response"
	"core/types"
	"core/utils"
	"core/web"
	"fmt"
	"strconv"
	"time"
)

type AdminNovelHandler struct {
	NovelDb *storage.NovelDb
}

func (e *AdminNovelHandler) GetNovelPage(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	searchKey := wc.Query("searchKey")
	var pagination = resp.NewPagination(int(pageNum), int(pageSize))
	records, total, err := e.NovelDb.GetNovelPage(searchKey, pagination.StartRow, pagination.EndRow)
	if err != nil {
		wc.AbortWithError(err)
	}
	pagination.SetRecords(records)
	pagination.SetTotal(total)
	return response.ReturnOK(pagination)
}

func (e *AdminNovelHandler) SaveNovel(wc *web.WebContext) interface{} {
	e.saveOrUpdateNovel(wc)
	return response.Success()
}

func (e *AdminNovelHandler) UpdateNovel(wc *web.WebContext) interface{} {
	e.saveOrUpdateNovel(wc)
	return response.Success()
}

func (e *AdminNovelHandler) saveOrUpdateNovel(wc *web.WebContext) {
	var params map[string]any
	if err := wc.ShouldBindJSON(&params); err != nil {
		wc.AbortWithError(err)
	}
	novelIdStr := fmt.Sprint(params["novelId"])
	firstLetter := fmt.Sprint(params["firstLetter"])
	title := fmt.Sprint(params["title"])
	cover := fmt.Sprint(params["cover"])
	authorIds := fmt.Sprint(params["authorIds"])
	authors := fmt.Sprint(params["authors"])
	categoryIds := fmt.Sprint(params["categoryIds"])
	categories := fmt.Sprint(params["categories"])
	regionIdStr := fmt.Sprint(params["regionId"])
	region := fmt.Sprint(params["region"])
	description := fmt.Sprint(params["description"])
	directionStr := fmt.Sprint(params["direction"])
	flagStr := fmt.Sprint(params["flag"])
	startTimeStr := fmt.Sprint(params["startTime"])
	endTimeStr := fmt.Sprint(params["endTime"])
	statusStr := fmt.Sprint(params["status"])
	var novel = new(model.Novel)
	if len(novelIdStr) > 0 {
		novelId, err := strconv.ParseInt(novelIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.NovelId = novelId
	}
	if len(regionIdStr) > 0 {
		regionId, err := strconv.ParseInt(regionIdStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.RegionId = regionId
	}
	if len(directionStr) > 0 {
		direction, err := strconv.ParseInt(directionStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.Direction = int8(direction)
	}
	if len(flagStr) > 0 {
		flag, err := strconv.ParseInt(flagStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.Flag = int8(flag)
	}
	if len(startTimeStr) > 0 {
		startTime, err := time.Parse(utils.NORM_DATETIME_PATTERN, startTimeStr)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.StartTime = types.Time(startTime)
	}
	if len(endTimeStr) > 0 {
		endTime, err := time.Parse(utils.NORM_DATETIME_PATTERN, endTimeStr)
		if err != nil {
			wc.AbortWithError(err)
		}
		novel.EndTime = types.Time(endTime)
	}
	status, err := strconv.ParseInt(statusStr, 10, 8)
	if err != nil {
		wc.AbortWithError(err)
	}
	novel.FirstLetter = firstLetter
	novel.Title = title
	novel.Cover = cover
	novel.AuthorIds = authorIds
	novel.Authors = authors
	novel.CategoryIds = categoryIds
	novel.Categories = categories
	novel.Region = region
	novel.Description = description
	novel.Status = int8(status)
	e.NovelDb.SaveOrUpdateNovel(novel)
}
