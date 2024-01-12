package resp

import "core/types"

type RankItemResp struct {
	ObjId      int64      `json:"objId"`
	Title      string     `json:"title"`
	Cover      string     `json:"cover"`
	AssetType  int8       `json:"assetType"`
	Categories string     `json:"categories"`
	Authors    string     `json:"authors"`
	Score      int64      `json:"score"`
	UpdatedAt  types.Time `json:"updatedAt"`
}
