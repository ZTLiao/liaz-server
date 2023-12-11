package handler

import (
	"business/enums"
	"business/resp"
	"business/storage"
	"core/constant"
	"core/response"
	"core/web"
	"strconv"
)

type RecommendHandler struct {
	RecommendDb     *storage.RecommendDb
	RecommendItemDb *storage.RecommendItemDb
	RecommendCache  *storage.RecommendCache
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
		recommendResp.RecommendId = recommend.RecommendId
		recommendResp.Title = recommend.Title
		recommendResp.ShowType = recommend.ShowType
		recommendResp.IsShowTitle = recommend.ShowTitle == constant.YES
		recommendResp.OptType = recommend.OptType
		recommendResp.OptValue = recommend.OptValue
		recommendResp.Items = recommendItemResps
		//设置缓存
		e.RecommendCache.HSet(int8(position), recommendType, recommendResp)
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
		//TODO
	} else if enums.RECOMMEND_TYPE_FOR_UPGRADE == recommendType {
		//TODO
	}
	return recommendItemResps, nil
}
