package handler

import (
	basicStorage "basic/storage"
	"business/enums"
	"business/model"
	"business/resp"
	businessStorage "business/storage"
	"core/response"
	"core/utils"
	"core/web"
	"sort"
	"strconv"
	"strings"
)

type NovelHandler struct {
	NovelDb            *businessStorage.NovelDb
	NovelVolumeDb      *businessStorage.NovelVolumeDb
	NovelChapterDb     *businessStorage.NovelChapterDb
	NovelChapterItemDb *businessStorage.NovelChapterItemDb
	FileItemDb         *basicStorage.FileItemDb
	NovelSubscribeDb   *businessStorage.NovelSubscribeDb
	BrowseDb           *businessStorage.BrowseDb
}

func (e *NovelHandler) NovelDetail(wc *web.WebContext) interface{} {
	novelIdStr := wc.Param("novelId")
	if len(novelIdStr) == 0 {
		return response.Success()
	}
	novelId, err := strconv.ParseInt(novelIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	novelDetail, err := e.GetNovelDetail(novelId)
	if err != nil {
		wc.AbortWithError(err)
	}
	if novelDetail == nil {
		return response.Success()
	}
	userId := web.GetUserId(wc)
	if userId != 0 {
		isSubscribe, err := e.NovelSubscribeDb.IsSubscribe(novelId, userId)
		if err != nil {
			wc.AbortWithError(err)
		}
		novelDetail.IsSubscribe = isSubscribe
		browse, err := e.BrowseDb.GetBrowseByObjId(userId, enums.ASSET_TYPE_FOR_NOVEL, novelId)
		if err != nil {
			wc.AbortWithError(err)
		}
		if browse != nil {
			novelDetail.BrowseChapterId = browse.ChapterId
			novelDetail.CurrentIndex = browse.StopIndex
		}
	}
	return response.ReturnOK(novelDetail)
}

func (e *NovelHandler) GetNovelDetail(novelId int64) (*resp.NovelDetailResp, error) {
	novel, err := e.NovelDb.GetNovelById(novelId)
	if err != nil {
		return nil, err
	}
	if novel == nil {
		return nil, nil
	}
	novelChapters, err := e.NovelChapterDb.GetNovelChapters(novelId)
	if err != nil {
		return nil, err
	}
	novelChapterItems, err := e.NovelChapterItemDb.GetNovelChapterItemByNovelId(novelId)
	if err != nil {
		return nil, err
	}
	var novelDetail = new(resp.NovelDetailResp)
	novelDetail.NovelId = novel.NovelId
	novelDetail.Title = novel.Title
	novelDetail.Cover = novel.Cover
	var authorIdList = make([]int64, 0)
	authorIds := strings.Split(novel.AuthorIds, utils.COMMA)
	for _, v := range authorIds {
		authorId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		authorIdList = append(authorIdList, authorId)
	}
	novelDetail.AuthorIds = authorIdList
	novelDetail.Authors = strings.Split(novel.Authors, utils.COMMA)
	var categoryIdList = make([]int64, 0)
	categoryIds := strings.Split(novel.CategoryIds, utils.COMMA)
	for _, v := range categoryIds {
		categoryId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		categoryIdList = append(categoryIdList, categoryId)
	}
	novelDetail.CategoryIds = categoryIdList
	novelDetail.Categories = strings.Split(novel.Categories, utils.COMMA)
	novelDetail.SubscribeNum = novel.SubscribeNum
	novelDetail.HitNum = novel.HitNum
	novelDetail.Updated = novel.EndTime
	novelDetail.Description = novel.Description
	novelDetail.Flag = novel.Flag
	novelDetail.Direction = novel.Direction
	if len(novelChapters) > 0 {
		var chapterItemMap = make(map[int64][]model.NovelChapterItem, 0)
		for _, novelChapterItem := range novelChapterItems {
			novelChapterId := novelChapterItem.NovelChapterId
			items, ex := chapterItemMap[novelChapterId]
			if !ex {
				items = make([]model.NovelChapterItem, 0)
			}
			items = append(items, novelChapterItem)
			chapterItemMap[novelChapterId] = items
		}
		var chapterMap = make(map[int64][]resp.NovelChapterResp, 0)
		for _, novelChapter := range novelChapters {
			novelChapterId := novelChapter.NovelChapterId
			novelVolumeId := novelChapter.NovelVolumeId
			chapters, ex := chapterMap[novelVolumeId]
			if !ex {
				chapters = make([]resp.NovelChapterResp, 0)
			}
			var paths = make([]string, 0)
			var types = make([]string, 0)
			items, ex := chapterItemMap[novelChapterId]
			if ex {
				sort.Slice(items, func(i, j int) bool {
					return items[i].SeqNo > items[j].SeqNo
				})
				for _, item := range items {
					path := item.Path
					fileType, err := e.FileItemDb.GetFileTypeByPath(path)
					if err != nil {
						return nil, err
					}
					paths = append(paths, path)
					types = append(types, fileType)
				}
			}
			var chapter = resp.NovelChapterResp{
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
			chapters = append(chapters, chapter)
			chapterMap[novelVolumeId] = chapters
		}
		var novelVolumes = make([]resp.NovelVolumeResp, 0)
		for k, v := range chapterMap {
			var volumeName string
			var seqNo int64
			if k != 0 {
				novelVolume, err := e.NovelVolumeDb.GetNovelVolumeById(k)
				if err != nil {
					return nil, err
				}
				if novelVolume != nil {
					volumeName = novelVolume.VolumeName
					seqNo = novelVolume.SeqNo
				}
			}
			novelVolumes = append(novelVolumes, resp.NovelVolumeResp{
				NovelVolumeId: k,
				VolumeName:    volumeName,
				SeqNo:         seqNo,
				Chapters:      v,
			})
		}
		sort.Slice(novelVolumes, func(i, j int) bool {
			return novelVolumes[i].SeqNo > novelVolumes[j].SeqNo
		})
		novelDetail.Volumes = novelVolumes
	}
	return novelDetail, nil
}

func (e *NovelHandler) NovelUpgrade(wc *web.WebContext) interface{} {
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	novels, err := e.NovelDb.GetNovelUpgradePage(int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(novels) == 0 {
		return response.Success()
	}
	var novelUpgrades = make([]resp.NovelUpgradeResp, 0)
	for _, novel := range novels {
		novelId := novel.NovelId
		novelChapter, err := e.NovelChapterDb.UpgradeChapter(novelId)
		if err != nil {
			wc.AbortWithError(err)
		}
		if novelChapter == nil {
			continue
		}
		var novelUpgrade = &resp.NovelUpgradeResp{
			NovelChapterId: novelChapter.NovelChapterId,
			NovelId:        novelId,
			Title:          novel.Title,
			Cover:          novel.Cover,
			Categories:     strings.Split(novel.Categories, utils.COMMA),
			Authors:        strings.Split(novel.Authors, utils.COMMA),
			UpgradeChapter: novelChapter.ChapterName,
			Updated:        novel.EndTime,
		}
		novelUpgrades = append(novelUpgrades, *novelUpgrade)
	}
	return response.ReturnOK(novelUpgrades)
}

func (e *NovelHandler) GetNovelByCategory(wc *web.WebContext) interface{} {
	categoryIdStr := wc.Query("categoryId")
	if len(categoryIdStr) == 0 {
		return response.Success()
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
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
	novels, err := e.NovelDb.GetNovelByCategory(categoryId, int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(novels) == 0 {
		return response.Success()
	}
	var categoryItems = make([]resp.CategoryItemResp, 0)
	for _, novel := range novels {
		novelId := novel.NovelId
		novelChapter, err := e.NovelChapterDb.UpgradeChapter(novelId)
		if err != nil {
			wc.AbortWithError(err)
		}
		var categoryItem = &resp.CategoryItemResp{
			CategoryId:     categoryId,
			AssetType:      enums.ASSET_TYPE_FOR_NOVEL,
			Title:          novel.Title,
			Cover:          novel.Cover,
			UpgradeChapter: novelChapter.ChapterName,
			ObjId:          novelChapter.NovelChapterId,
		}
		categoryItems = append(categoryItems, *categoryItem)
	}
	return response.ReturnOK(categoryItems)
}

func (e *NovelHandler) NovelCatalogue(wc *web.WebContext) interface{} {
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
	novelDetail, err := e.GetNovelDetail(novelChapter.NovelId)
	if err != nil {
		wc.AbortWithError(err)
	}
	volumes := novelDetail.Volumes
	if len(volumes) == 0 {
		return response.Success()
	}
	var chapters []resp.NovelChapterResp
	for _, v := range volumes {
		if v.NovelVolumeId == novelChapter.NovelVolumeId {
			chapters = v.Chapters
			break
		}
	}
	return response.ReturnOK(chapters)
}

func (e *NovelHandler) GetNovel(wc *web.WebContext) interface{} {
	novelIdStr := wc.Query("novelId")
	if len(novelIdStr) == 0 {
		return response.Success()
	}
	novelId, err := strconv.ParseInt(novelIdStr, 10, 64)
	if err != nil {
		wc.AbortWithError(err)
	}
	novel, err := e.NovelDb.GetNovelById(novelId)
	if err != nil {
		wc.AbortWithError(err)
	}
	return response.ReturnOK(novel)
}
