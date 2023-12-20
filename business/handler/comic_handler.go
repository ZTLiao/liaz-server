package handler

import (
	"business/model"
	"business/resp"
	"business/storage"
	"core/response"
	"core/utils"
	"core/web"
	"sort"
	"strconv"
	"strings"
)

type ComicHandler struct {
	ComicDb            *storage.ComicDb
	ComicChapterDb     *storage.ComicChapterDb
	ComicChapterItemDb *storage.ComicChapterItemDb
}

func (e *ComicHandler) ComicDetail(wc *web.WebContext) interface{} {
	comicIdStr := wc.Param("comicId")
	if len(comicIdStr) == 0 {
		return response.Success()
	}
	comicId, err := strconv.ParseInt(comicIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	comic, err := e.ComicDb.GetComicById(comicId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if comic == nil {
		return response.Success()
	}
	comicChapters, err := e.ComicChapterDb.GetComicChapters(comicId)
	if err != nil {
		wc.AbortWithError(err)
	}
	comicChapterItems, err := e.ComicChapterItemDb.GetComicChapterItems(comicId)
	if err != nil {
		wc.AbortWithError(err)
	}
	var comicDetail = new(resp.ComicDetailResp)
	comicDetail.ComicId = comic.ComicId
	comicDetail.Title = comic.Title
	comicDetail.Cover = comic.Cover
	var authorIdList = make([]int64, 0)
	authorIds := strings.Split(comic.AuthorIds, utils.COMMA)
	for _, v := range authorIds {
		authorId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		authorIdList = append(authorIdList, authorId)
	}
	comicDetail.AuthorIds = authorIdList
	comicDetail.Authors = strings.Split(comic.Authors, utils.COMMA)
	var categoryIdList = make([]int64, 0)
	categoryIds := strings.Split(comic.CategoryIds, utils.COMMA)
	for _, v := range categoryIds {
		categoryId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
		categoryIdList = append(categoryIdList, categoryId)
	}
	comicDetail.CategoryIds = categoryIdList
	comicDetail.Categories = strings.Split(comic.Categories, utils.COMMA)
	comicDetail.SubscribeNum = comic.SubscribeNum
	comicDetail.HitNum = comic.HitNum
	comicDetail.Updated = comic.EndTime
	comicDetail.Description = comic.Description
	comicDetail.Flag = comic.Flag
	comicDetail.Direction = comic.Direction
	if len(comicChapters) > 0 {
		var chapterItemMap = make(map[int64][]model.ComicChapterItem, 0)
		for _, comicChapterItem := range comicChapterItems {
			comicChapterId := comicChapterItem.ComicChapterId
			items, ex := chapterItemMap[comicChapterId]
			if !ex {
				items = make([]model.ComicChapterItem, 0)
			}
			items = append(items, comicChapterItem)
			chapterItemMap[comicChapterId] = items
		}
		var chapterMap = make(map[int8][]resp.ComicChapterResp, 0)
		for _, comicChapter := range comicChapters {
			comicChapterId := comicChapter.ComicChapterId
			chapterType := comicChapter.ChapterType
			chapters, ex := chapterMap[chapterType]
			if !ex {
				chapters = make([]resp.ComicChapterResp, 0)
			}
			var paths = make([]string, 0)
			items, ex := chapterItemMap[comicChapterId]
			if ex {
				sort.Slice(items, func(i, j int) bool {
					return items[i].SeqNo < items[j].SeqNo
				})
				for _, item := range items {
					paths = append(paths, item.Path)
				}
			}
			var chapter = resp.ComicChapterResp{
				ComicChapterId: comicChapter.ComicChapterId,
				ComicId:        comicChapter.ComicId,
				Flag:           comic.Flag,
				ChapterName:    comicChapter.ChapterName,
				ChapterType:    comicChapter.ChapterType,
				PageNum:        len(paths),
				SeqNo:          int(comicChapter.SeqNo),
				Direction:      comic.Direction,
				UpdatedAt:      comicChapter.UpdatedAt,
				Paths:          paths,
			}
			chapters = append(chapters, chapter)
			chapterMap[chapterType] = chapters
		}
		var chapterTypes []resp.ComicChapterTypeResp
		for k, v := range chapterMap {
			chapterTypes = append(chapterTypes, resp.ComicChapterTypeResp{
				ChapterType: k,
				Flag:        comic.Flag,
				Chapters:    v,
			})
		}
		comicDetail.ChapterTypes = chapterTypes
	}
	return response.ReturnOK(comicDetail)
}

func (e *ComicHandler) ComicUpgrade(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	comics, err := e.ComicDb.GetComicUpgradePage(int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(comics) == 0 {
		return response.Success()
	}
	var comicUpgrades = make([]resp.ComicUpgradeResp, 0)
	for _, comic := range comics {
		comicId := comic.ComicId
		comicChapter, err := e.ComicChapterDb.UpgradeChapter(comicId)
		if err != nil {
			wc.AbortWithError(err)
		}
		if comicChapter == nil {
			continue
		}
		var comicUpgrade = &resp.ComicUpgradeResp{
			ComicChapterId: comicChapter.ComicChapterId,
			ComicId:        comicId,
			Title:          comic.Title,
			Cover:          comic.Cover,
			Categories:     strings.Split(comic.Categories, utils.COMMA),
			Authors:        strings.Split(comic.Authors, utils.COMMA),
			UpgradeChapter: comicChapter.ChapterName,
			Updated:        comic.EndTime,
		}
		comicUpgrades = append(comicUpgrades, *comicUpgrade)
	}
	return response.ReturnOK(comicUpgrades)
}
