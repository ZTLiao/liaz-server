package handler

import (
	"business/resp"
	"business/storage"
	"core/response"
	"core/utils"
	"core/web"
	"strconv"
	"strings"
)

type ComicHandler struct {
	ComicDb            *storage.ComicDb
	ComicChapterDb     *storage.ComicChapterDb
	ComicChapterItemDb *storage.ComicChapterItemDb
}

func (e *ComicHandler) ComicDetail(wc *web.WebContext) interface{} {
	return response.ReturnOK(nil)
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
