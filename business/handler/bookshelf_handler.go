package handler

import (
	"business/enums"
	"business/resp"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type BookshelfHandler struct {
	ComicChapterDb *storage.ComicChapterDb
	NovelChapterDb *storage.NovelChapterDb
}

func (e *BookshelfHandler) GetComic(wc *web.WebContext) interface{} {
	sortType, err := strconv.ParseInt(wc.DefaultQuery("sortType", "0"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	if userId == 0 {
		return response.Success()
	}
	comicChapters, err := e.ComicChapterDb.GetBookshelf(userId, int32(sortType), int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(comicChapters) == 0 {
		return response.Success()
	}
	var categoryItems = make([]resp.CategoryItemResp, 0)
	for _, v := range comicChapters {
		categoryItems = append(categoryItems, resp.CategoryItemResp{
			CategoryId:     v.ComicId,
			AssetType:      enums.ASSET_TYPE_FOR_COMIC,
			Title:          v.Title,
			Cover:          v.Cover,
			ObjId:          v.ChapterId,
			UpgradeChapter: v.ChapterName,
			UpdatedAt:      v.EndTime,
		})
	}
	return response.ReturnOK(categoryItems)
}

func (e *BookshelfHandler) GetNovel(wc *web.WebContext) interface{} {
	sortType, err := strconv.ParseInt(wc.DefaultQuery("sortType", "0"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	userId := web.GetUserId(wc)
	if userId == 0 {
		return response.Success()
	}
	novelChapters, err := e.NovelChapterDb.GetBookshelf(userId, int32(sortType), int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(novelChapters) == 0 {
		return response.Success()
	}
	var categoryItems = make([]resp.CategoryItemResp, 0)
	for _, v := range novelChapters {
		categoryItems = append(categoryItems, resp.CategoryItemResp{
			CategoryId:     v.NovelId,
			AssetType:      enums.ASSET_TYPE_FOR_NOVEL,
			Title:          v.Title,
			Cover:          v.Cover,
			ObjId:          v.ChapterId,
			UpgradeChapter: v.ChapterName,
			UpdatedAt:      v.EndTime,
		})
	}
	return response.ReturnOK(categoryItems)
}
