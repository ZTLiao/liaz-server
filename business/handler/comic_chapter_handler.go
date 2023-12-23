package handler

import (
	"business/resp"
	"business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type ComicChapterHandler struct {
	ComicDb            *storage.ComicDb
	ComicChapterDb     *storage.ComicChapterDb
	ComicChapterItemDb *storage.ComicChapterItemDb
}

func (e *ComicChapterHandler) GetComicChapter(wc *web.WebContext) interface{} {
	comicChapterIdStr := wc.Query("comicChapterId")
	if len(comicChapterIdStr) == 0 {
		return response.Success()
	}
	comicChapterId, err := strconv.ParseInt(comicChapterIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	comicChapter, err := e.ComicChapterDb.GetComicChapterById(comicChapterId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if comicChapter == nil {
		return response.Success()
	}
	paths, err := e.ComicChapterItemDb.GetPathByComicChapterId(comicChapterId)
	if err != nil {
		wc.AbortWithError(err)
	}
	comic, err := e.ComicDb.GetComicById(comicChapter.ComicId)
	if err != nil {
		wc.AbortWithError(err)
	}
	var comicChapterResp = &resp.ComicChapterResp{
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
	return response.ReturnOK(comicChapterResp)
}
