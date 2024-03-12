package handler

import (
	"basic/handler"
	"business/enums"
	"business/model"
	"business/resp"
	"business/storage"
	"core/constant"
	"core/response"
	"core/utils"
	"core/web"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type RecommendHandler struct {
	ComicDb         *storage.ComicDb
	NovelDb         *storage.NovelDb
	RecommendDb     *storage.RecommendDb
	RecommendItemDb *storage.RecommendItemDb
	RecommendCache  *storage.RecommendCache
	AssetDb         *storage.AssetDb
	SysConfHandler  *handler.SysConfHandler
	ComicRankCache  *storage.ComicRankCache
	NovelRankCache  *storage.NovelRankCache
}

func (e *RecommendHandler) Recommend(wc *web.WebContext) interface{} {
	positionStr := wc.Param("position")
	var position int64
	var err error
	if len(positionStr) > 0 {
		position, err = strconv.ParseInt(positionStr, 10, 64)
		if err != nil {
			wc.AbortWithError(err)
		}
	}
	var recommendResps []resp.RecommendResp
	recommendResps, err = e.RecommendCache.HGetAll(int8(position))
	if err != nil {
		wc.AbortWithError(err)
	}
	if len(recommendResps) == 0 {
		recommendResps, err = e.GetRecommend(int8(position))
		if err != nil {
			wc.AbortWithError(err)
		}
	} else {
		go e.GetRecommend(int8(position))
	}
	userId := web.GetUserId(wc)
	if userId != 0 && len(recommendResps) > 0 {
		for i, recommendResp := range recommendResps {
			recommendType := recommendResp.RecommendType
			if enums.RECOMMEND_TYPE_FOR_MY_SUBSCRIBE == recommendType {
				intValue, err := e.SysConfHandler.GetIntValueByKey(constant.RECOMMEND_FOR_MY_SUBSCRIBE)
				if err != nil {
					wc.AbortWithError(err)
				}
				if intValue == 0 {
					intValue = 20
				}
				assets, err := e.AssetDb.GetAssetForMySubscribe(userId, int64(intValue))
				if err != nil {
					wc.AbortWithError(err)
				}
				recommendResp.Items = e.convertItem(assets, true)
				recommendResps[i] = recommendResp
				break
			}
		}
	}
	sort.Slice(recommendResps, func(i, j int) bool {
		return recommendResps[i].SeqNo < recommendResps[j].SeqNo
	})
	return response.ReturnOK(recommendResps)
}

func (e *RecommendHandler) GetRecommend(position int8) ([]resp.RecommendResp, error) {
	var recommendResps []resp.RecommendResp
	recommends, err := e.RecommendDb.GetRecommendByPosition(position)
	if err != nil {
		return nil, err
	}
	for _, recommend := range recommends {
		recommendId := recommend.RecommendId
		recommendType := recommend.RecommendType
		recommendItemResps, err := e.GetRecommendItems(recommendId, recommendType)
		if err != nil {
			return nil, err
		}
		var recommendResp = new(resp.RecommendResp)
		recommendResp.RecommendId = recommendId
		recommendResp.RecommendType = recommendType
		recommendResp.Title = recommend.Title
		recommendResp.ShowType = recommend.ShowType
		recommendResp.IsShowTitle = recommend.ShowTitle == constant.YES
		recommendResp.OptType = recommend.OptType
		recommendResp.OptValue = recommend.OptValue
		recommendResp.SeqNo = recommend.SeqNo
		recommendResp.Items = recommendItemResps
		//设置缓存
		if recommendResp.RecommendId != 0 {
			e.RecommendCache.HSet(int8(position), recommendType, recommendResp)
		}
		recommendResps = append(recommendResps, *recommendResp)
	}
	return recommendResps, nil
}

func (e *RecommendHandler) GetRecommendItems(recommendId int64, recommendType int8) ([]resp.RecommendItemResp, error) {
	var recommendItemResps []resp.RecommendItemResp
	if enums.RECOMMEND_TYPE_FOR_CUSTOM == recommendType {
		recommendItems, err := e.RecommendItemDb.GetRecommendItemByRecommendId(recommendId)
		if err != nil {
			return nil, err
		}
		for _, recommendItem := range recommendItems {
			recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
				RecommendItemId: recommendItem.RecommendItemId,
				Title:           recommendItem.Title,
				SubTitle:        recommendItem.SubTitle,
				ShowValue:       recommendItem.ShowValue,
				SkipType:        recommendItem.SkipType,
				SkipValue:       recommendItem.SkipValue,
				ObjId:           recommendItem.ObjId,
			})
		}
	} else if enums.RECOMMEND_TYPE_FOR_HOT == recommendType {
		intValue, err := e.SysConfHandler.GetIntValueByKey(constant.RECOMMEND_FOR_HOT)
		if err != nil {
			return nil, err
		}
		if intValue == 0 {
			intValue = 9
		}
		var now = time.Now()
		dateTime := utils.GetStartOfWeek(now).Format(utils.NORM_DATE_PATTERN) + utils.COLON + utils.GetEndOfWeek(now).Format(utils.NORM_DATE_PATTERN)
		comicRankMap, err := e.ComicRankCache.Rank(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, dateTime, 0, 30)
		if err != nil {
			return nil, err
		}
		var comicIds = make([]int64, 0)
		for k := range comicRankMap {
			comicIds = append(comicIds, k)
		}
		novelRankMap, err := e.NovelRankCache.Rank(enums.RANK_TYPE_FOR_POPULAR, enums.TIME_TYPE_FOR_WEEK, dateTime, 0, 30)
		if err != nil {
			return nil, err
		}
		var novelIds = make([]int64, 0)
		for k := range novelRankMap {
			novelIds = append(novelIds, k)
		}
		var assets = make([]model.Asset, 0)
		comicAssets, err := e.AssetDb.GetAssetByObjId(comicIds, enums.ASSET_TYPE_FOR_COMIC)
		if err != nil {
			return nil, err
		}
		if len(comicAssets) != 0 {
			assets = append(assets, comicAssets...)
		}
		novelAssets, err := e.AssetDb.GetAssetByObjId(novelIds, enums.ASSET_TYPE_FOR_NOVEL)
		if err != nil {
			return nil, err
		}
		if len(novelAssets) != 0 {
			assets = append(assets, novelAssets...)
		}
		rand.New(rand.NewSource(now.UnixNano()))
		var recommendAssets = make([]model.Asset, 0)
		if len(assets) > intValue {
			var index int
			var set = make(map[int]bool)
			for i := 0; i < int(intValue); i++ {
				index = rand.Intn(len(assets))
				for ok := set[index]; ok; {
					index = rand.Intn(len(assets))
				}
				set[index] = true
				recommendAssets = append(recommendAssets, assets[index])
			}
		} else {
			recommendAssets = append(recommendAssets, assets...)
		}
		recommendItemResps = e.convertItem(recommendAssets, false)
	} else if enums.RECOMMEND_TYPE_FOR_UPGRADE == recommendType {
		intValue, err := e.SysConfHandler.GetIntValueByKey(constant.RECOMMEND_FOR_UPGRADE)
		if err != nil {
			return nil, err
		}
		if intValue == 0 {
			intValue = 9
		}
		assets, err := e.AssetDb.GetAssetForUpgrade(int64(intValue))
		if err != nil {
			return nil, err
		}
		recommendItemResps = e.convertItem(assets, false)
	}
	return recommendItemResps, nil
}

func (e *RecommendHandler) convertItem(assets []model.Asset, isUpgrade bool) []resp.RecommendItemResp {
	var recommendItemResps []resp.RecommendItemResp
	if len(assets) == 0 {
		return recommendItemResps
	}
	for _, asset := range assets {
		subTitle := asset.UpgradeChapter
		assetKey := asset.AssetKey
		if !isUpgrade && len(assetKey) > 0 {
			array := strings.Split(assetKey, utils.PIPE)
			if len(array) > 1 {
				subTitle = array[1]
			}
		}
		recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
			RecommendItemId: asset.AssetId,
			Title:           asset.Title,
			SubTitle:        subTitle,
			ShowValue:       asset.Cover,
			SkipType:        asset.AssetType,
			SkipValue:       strconv.FormatInt(asset.ChapterId, 10),
			ObjId:           strconv.FormatInt(asset.ObjId, 10),
		})
	}
	return recommendItemResps
}

func (e *RecommendHandler) DelRecommendCache(recommendId int64) error {
	recommend, err := e.RecommendDb.GetRecommendById(recommendId)
	if err != nil {
		return err
	}
	if recommend == nil {
		return nil
	}
	position := recommend.Position
	err = e.RecommendCache.Del(position)
	if err != nil {
		return err
	}
	return nil
}

func (e *RecommendHandler) RecommendComic(wc *web.WebContext) interface{} {
	comicIdStr := wc.Query("comicId")
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
	var comicIdMap = make(map[int64]int64, 0)
	var recommendResps []resp.RecommendResp
	authorIds := comic.AuthorIds
	authors := comic.Authors
	if len(authorIds) > 0 {
		authorIdArray := strings.Split(authorIds, utils.COMMA)
		authorArray := strings.Split(authors, utils.COMMA)
		for i, authorIdStr := range authorIdArray {
			author := authorArray[i]
			var recommendResp = new(resp.RecommendResp)
			recommendResp.RecommendId = comicId
			recommendResp.RecommendType = enums.RECOMMEND_TYPE_FOR_AUTHOR
			recommendResp.Title = author
			recommendResp.ShowType = enums.SHOW_TYPE_FOR_NONE
			recommendResp.IsShowTitle = true
			recommendResp.OptType = enums.OPT_TYPE_FOR_JUMP
			recommendResp.OptValue = authorIdStr
			recommendResp.SeqNo = i
			authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			var recommendItemResps []resp.RecommendItemResp
			novels, err := e.NovelDb.GetNovelByAuthor(authorId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(novels) > 0 {
				for _, novel := range novels {
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: novel.NovelId,
						Title:           novel.Title,
						SubTitle:        novel.Authors,
						ShowValue:       novel.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_NOVEL,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(novel.NovelId, 10),
					})
				}
				recommendResp.Items = recommendItemResps
			}
			comics, err := e.ComicDb.GetComicByAuthor(authorId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(comics) > 0 {
				for _, comic := range comics {
					if comicId == comic.ComicId {
						continue
					}
					if _, ex := comicIdMap[comic.ComicId]; ex {
						continue
					}
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: comic.ComicId,
						Title:           comic.Title,
						SubTitle:        comic.Authors,
						ShowValue:       comic.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_COMIC,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(comic.ComicId, 10),
					})
					comicIdMap[comic.ComicId] = comic.ComicId
				}
				recommendResp.Items = recommendItemResps
			}
			recommendResps = append(recommendResps, *recommendResp)
		}
	}
	categoryIds := comic.CategoryIds
	categories := comic.Categories
	if len(categoryIds) > 0 {
		categoryIdArray := strings.Split(categoryIds, utils.COMMA)
		categoryArray := strings.Split(categories, utils.COMMA)
		for i, categoryIdStr := range categoryIdArray {
			category := categoryArray[i]
			var recommendResp = new(resp.RecommendResp)
			recommendResp.RecommendId = comicId
			recommendResp.RecommendType = enums.RECOMMEND_TYPE_FOR_CATEGORY
			recommendResp.Title = category
			recommendResp.ShowType = enums.SHOW_TYPE_FOR_NONE
			recommendResp.IsShowTitle = true
			recommendResp.OptType = enums.OPT_TYPE_FOR_JUMP
			recommendResp.OptValue = categoryIdStr
			recommendResp.SeqNo = i
			categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			var recommendItemResps []resp.RecommendItemResp
			comics, err := e.ComicDb.GetComicByCategory(categoryId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(comics) > 0 {
				for _, comic := range comics {
					if comicId == comic.ComicId {
						continue
					}
					if _, ex := comicIdMap[comic.ComicId]; ex {
						continue
					}
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: comic.ComicId,
						Title:           comic.Title,
						SubTitle:        comic.Authors,
						ShowValue:       comic.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_COMIC,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(comic.ComicId, 10),
					})
					comicIdMap[comic.ComicId] = comic.ComicId
				}
				recommendResp.Items = recommendItemResps
			}
			recommendResps = append(recommendResps, *recommendResp)
		}
	}
	return response.ReturnOK(recommendResps)
}

func (e *RecommendHandler) RecommendNovel(wc *web.WebContext) interface{} {
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
	if novel == nil {
		return response.Success()
	}
	var novelIdMap = make(map[int64]int64, 0)
	var recommendResps []resp.RecommendResp
	authorIds := novel.AuthorIds
	authors := novel.Authors
	if len(authorIds) > 0 {
		authorIdArray := strings.Split(authorIds, utils.COMMA)
		authorArray := strings.Split(authors, utils.COMMA)
		for i, authorIdStr := range authorIdArray {
			author := authorArray[i]
			var recommendResp = new(resp.RecommendResp)
			recommendResp.RecommendId = novelId
			recommendResp.RecommendType = enums.RECOMMEND_TYPE_FOR_AUTHOR
			recommendResp.Title = author
			recommendResp.ShowType = enums.SHOW_TYPE_FOR_NONE
			recommendResp.IsShowTitle = true
			recommendResp.OptType = enums.OPT_TYPE_FOR_JUMP
			recommendResp.OptValue = authorIdStr
			recommendResp.SeqNo = i
			authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			var recommendItemResps []resp.RecommendItemResp
			comics, err := e.ComicDb.GetComicByAuthor(authorId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(comics) > 0 {
				for _, comic := range comics {
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: comic.ComicId,
						Title:           comic.Title,
						SubTitle:        comic.Authors,
						ShowValue:       comic.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_COMIC,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(comic.ComicId, 10),
					})
				}
				recommendResp.Items = recommendItemResps
			}
			novels, err := e.NovelDb.GetNovelByAuthor(authorId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(novels) > 0 {
				for _, novel := range novels {
					if novelId == novel.NovelId {
						continue
					}
					if _, ex := novelIdMap[novel.NovelId]; ex {
						continue
					}
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: novel.NovelId,
						Title:           novel.Title,
						SubTitle:        novel.Authors,
						ShowValue:       novel.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_NOVEL,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(novel.NovelId, 10),
					})
					novelIdMap[novel.NovelId] = novel.NovelId
				}
				recommendResp.Items = recommendItemResps
			}

			recommendResps = append(recommendResps, *recommendResp)
		}
	}
	categoryIds := novel.CategoryIds
	categories := novel.Categories
	if len(categoryIds) > 0 {
		categoryIdArray := strings.Split(categoryIds, utils.COMMA)
		categoryArray := strings.Split(categories, utils.COMMA)
		for i, categoryIdStr := range categoryIdArray {
			category := categoryArray[i]
			var recommendResp = new(resp.RecommendResp)
			recommendResp.RecommendId = novelId
			recommendResp.RecommendType = enums.RECOMMEND_TYPE_FOR_CATEGORY
			recommendResp.Title = category
			recommendResp.ShowType = enums.SHOW_TYPE_FOR_NONE
			recommendResp.IsShowTitle = true
			recommendResp.OptType = enums.OPT_TYPE_FOR_JUMP
			recommendResp.OptValue = categoryIdStr
			recommendResp.SeqNo = i
			categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
			if err != nil {
				wc.AbortWithError(err)
			}
			var recommendItemResps []resp.RecommendItemResp
			novels, err := e.NovelDb.GetNovelByCategory(categoryId, 1, 10)
			if err != nil {
				wc.AbortWithError(err)
			}
			if len(novels) > 0 {
				for _, novel := range novels {
					if novelId == novel.NovelId {
						continue
					}
					if _, ex := novelIdMap[novel.NovelId]; ex {
						continue
					}
					recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
						RecommendItemId: novel.NovelId,
						Title:           novel.Title,
						SubTitle:        novel.Authors,
						ShowValue:       novel.Cover,
						SkipType:        enums.SKIP_TYPE_FOR_NOVEL,
						SkipValue:       utils.EMPTY,
						ObjId:           strconv.FormatInt(novel.NovelId, 10),
					})
					novelIdMap[novel.NovelId] = novel.NovelId
				}
				recommendResp.Items = recommendItemResps
			}
			recommendResps = append(recommendResps, *recommendResp)
		}
	}
	return response.ReturnOK(recommendResps)
}
