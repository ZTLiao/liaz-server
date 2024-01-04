package handler

import (
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
		for _, recommendResp := range recommendResps {
			recommendType := recommendResp.RecommendType
			if enums.RECOMMEND_TYPE_FOR_MY_SUBSCRIBE == recommendType {
				var recommendItemResps []resp.RecommendItemResp
				assets, err := e.AssetDb.GetAssetForMySubscribe(userId, 18)
				if err != nil {
					wc.AbortWithError(err)
				}
				e.ConvertRecommendItem(assets, recommendItemResps)
				recommendResp.Items = recommendItemResps
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
		assets, err := e.AssetDb.GetAssetForHot(9)
		if err != nil {
			return nil, err
		}
		e.ConvertRecommendItem(assets, recommendItemResps)
	} else if enums.RECOMMEND_TYPE_FOR_UPGRADE == recommendType {
		assets, err := e.AssetDb.GetAssetForUpgrade(9)
		if err != nil {
			return nil, err
		}
		e.ConvertRecommendItem(assets, recommendItemResps)
	}
	return recommendItemResps, nil
}

func (e *RecommendHandler) ConvertRecommendItem(assets []model.Asset, recommendItemResps []resp.RecommendItemResp) {
	for _, asset := range assets {
		assetKey := asset.AssetKey
		var authors string
		if len(assetKey) > 0 {
			assetKeyArray := strings.Split(assetKey, utils.PIPE)
			if len(assetKeyArray) > 1 {
				authors = assetKeyArray[1]
			}
		}
		recommendItemResps = append(recommendItemResps, resp.RecommendItemResp{
			RecommendItemId: asset.AssetId,
			Title:           asset.Title,
			SubTitle:        authors,
			ShowValue:       asset.Cover,
			SkipType:        asset.AssetType,
			SkipValue:       strconv.FormatInt(asset.ObjId, 10),
			ObjId:           strconv.FormatInt(asset.ObjId, 10),
		})
	}
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
