package handler

import (
	basicStorage "basic/storage"
	"business/resp"
	businessStorage "business/storage"
	"core/response"
	"core/web"
	"strconv"
)

type NovelChapterHandler struct {
	NovelDb            *businessStorage.NovelDb
	NovelChapterDb     *businessStorage.NovelChapterDb
	NovelChapterItemDb *businessStorage.NovelChapterItemDb
	FileItemDb         *basicStorage.FileItemDb
}

func (e *NovelChapterHandler) GetNovelChapter(wc *web.WebContext) interface{} {
	novelChapterIdStr := wc.Query("novelChapterId")
	if len(novelChapterIdStr) == 0 {
		return response.Success()
	}
	novelChapterId, err := strconv.ParseInt(novelChapterIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	novelChapter, err := e.NovelChapterDb.GetNovelChapterById(novelChapterId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if novelChapter == nil {
		return response.Success()
	}
	paths, err := e.NovelChapterItemDb.GetPathByNovelChapterId(novelChapterId)
	if err != nil {
		wc.AbortWithError(err)
	}
	novel, err := e.NovelDb.GetNovelById(novelChapter.NovelId)
	if err != nil {
		wc.AbortWithError(err)
	}
	var types = make([]string, 0)
	if len(paths) > 0 {
		for _, path := range paths {
			fileType, err := e.FileItemDb.GetFileTypeByPath(path)
			if err != nil {
				wc.AbortWithError(err)
			}
			types = append(types, fileType)
		}
	}
	var novelChapterResp = &resp.NovelChapterResp{
		NovelChapterId: novelChapter.NovelChapterId,
		NovelId:        novelChapter.NovelId,
		Flag:           novel.Flag,
		ChapterName:    novelChapter.ChapterName,
		ChapterType:    novelChapter.ChapterType,
		PageNum:        len(paths),
		SeqNo:          int(novelChapter.SeqNo),
		Direction:      novel.Direction,
		UpdatedAt:      novelChapter.UpdatedAt,
		Paths:          paths,
		Types:          types,
	}
	return response.ReturnOK(novelChapterResp)
}
