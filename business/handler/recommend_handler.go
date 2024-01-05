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
	"sort"
	"strconv"
	"strings"
)

type RecommendHandler struct {
	RecommendDb     *storage.RecommendDb
	RecommendItemDb *storage.RecommendItemDb
	RecommendCache  *storage.RecommendCache
	AssetDb         *storage.AssetDb
	SysConfHandler  *handler.SysConfHandler
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
		assets, err := e.AssetDb.GetAssetForHot(int64(intValue))
		if err != nil {
			return nil, err
		}
		recommendItemResps = e.convertItem(assets, false)
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
			SkipValue:       strconv.FormatInt(asset.ObjId, 10),
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
