package handler

import (
	"basic/device"
	basicHandler "basic/handler"
	"business/enums"
	"business/resp"
	businessStorage "business/storage"
	"core/constant"
	"core/logger"
	"core/response"
	"core/utils"
	"core/web"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type SearchHandler struct {
	SysConfHandler        *basicHandler.SysConfHandler
	SearchDb              *businessStorage.SearchDb
	AssetDb               *businessStorage.AssetDb
	SearchCache           *businessStorage.SearchCache
	ComicUpgradeItemCache *businessStorage.ComicUpgradeItemCache
	NovelUpgradeItemCache *businessStorage.NovelUpgradeItemCache
}

func (e *SearchHandler) Search(wc *web.WebContext) interface{} {
	key := wc.Query("key")
	if len(key) == 0 {
		return response.Success()
	}
	pageNum, err := strconv.ParseInt(wc.DefaultQuery("pageNum", "1"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	pageSize, err := strconv.ParseInt(wc.DefaultQuery("pageSize", "10"), 10, 32)
	if err != nil {
		wc.AbortWithError(err)
	}
	searchs, err := e.AssetDb.Search(key, int32(pageNum), int32(pageSize))
	if err != nil {
		wc.AbortWithError(err)
	}
	var searchResps = make([]resp.SearchResp, 0)
	var result string
	if len(searchs) > 0 {
		var comicIds = make([]int64, 0)
		var novelIds = make([]int64, 0)
		var assetIds = make([]string, 0)
		for _, search := range searchs {
			assetId := search.AssetId
			objId := search.ObjId
			assetType := search.AssetType
			searchResps = append(searchResps, resp.SearchResp{
				ObjId:          objId,
				Title:          search.Title,
				Cover:          search.Cover,
				AssetType:      assetType,
				Categories:     search.Categories,
				Authors:        search.Authors,
				UpgradeChapter: search.UpgradeChapter,
			})
			if enums.ASSET_TYPE_FOR_COMIC == assetType {
				comicIds = append(comicIds, objId)
			} else if enums.ASSET_TYPE_FOR_NOVEL == assetType {
				novelIds = append(novelIds, objId)
			}
			assetIds = append(assetIds, strconv.FormatInt(assetId, 10))
			e.SearchCache.Incr(assetId)
		}
		result = strings.Join(assetIds, utils.COMMA)
		if pageNum == 1 {
			if len(comicIds) == 0 {
				e.AutoAddComicSearchJob(key)
			}
			if len(novelIds) == 0 {
				e.AutoAddNovelSearchJob(key)
			}
		}
	} else {
		if pageNum == 1 {
			go e.AutoAddSearchJob(key)
		}
	}
	deviceInfo := device.GetDeviceInfo(wc)
	userId := web.GetUserId(wc)
	e.SearchDb.SaveSearch(key, deviceInfo.DeviceId, userId, result)
	if len(searchResps) == 0 {
		message := fmt.Sprintf(constant.ADD_SEARCH_JOB_ERROR, key)
		return response.Fail(message)
	}
	return response.ReturnOK(searchResps)
}

func (e *SearchHandler) AutoAddSearchJob(key string) {
	e.AutoAddComicSearchJob(key)
	e.AutoAddNovelSearchJob(key)
}

func (e *SearchHandler) AutoAddComicSearchJob(key string) {
	searchKey := url.QueryEscape(key)
	comicSpider, err := e.SysConfHandler.GetConfValueByKey(constant.COMIC_SPIDER)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if len(comicSpider) != 0 {
		go func() {
			url := fmt.Sprintf(comicSpider, searchKey)
			_, err := http.Get(url)
			if err != nil {
				logger.Error(err.Error())
				return
			}
		}()
	}
	e.ComicUpgradeItemCache.Del()
}

func (e *SearchHandler) AutoAddNovelSearchJob(key string) {
	searchKey := url.QueryEscape(key)
	novelSpider, err := e.SysConfHandler.GetConfValueByKey(constant.NOVEL_SPIDER)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if len(novelSpider) != 0 {
		go func() {
			url := fmt.Sprintf(novelSpider, searchKey)
			_, err := http.Get(url)
			if err != nil {
				logger.Error(err.Error())
				return
			}
		}()
	}
	e.NovelUpgradeItemCache.Del()
}

func (e *SearchHandler) HotRank(wc *web.WebContext) interface{} {
	members, err := e.SearchCache.Rank(0, 30)
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(members) == 0 {
		return response.Success()
	}
	var assetIds = make([]int64, 0)
	for k := range members {
		assetIds = append(assetIds, k)
	}
	assets, err := e.AssetDb.GetAssetByIds(assetIds)
	if err != nil {
		wc.AbortWithError(err)
	}
	sort.Slice(assets, func(i, j int) bool {
		return members[assets[i].AssetId] > members[assets[j].AssetId]
	})
	return response.ReturnOK(assets)
}
