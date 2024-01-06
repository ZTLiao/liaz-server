package resp

type SearchResp struct {
	ObjId      int64  `json:"objId"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	AssetType  int8   `json:"assetType"`
	Categories string `json:"categories"`
	Authors    string `json:"authors"`
}
