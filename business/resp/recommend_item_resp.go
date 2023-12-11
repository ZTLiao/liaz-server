package resp

type RecommendItemResp struct {
	RecommendItemId int64  `json:"recommendItemId"`
	Title           string `json:"title"`
	SubTitle        string `json:"subTitle"`
	ShowValue       string `json:"showValue"`
	SkipType        int8   `json:"skipType"`
	SkipValue       string `json:"skipValue"`
	ObjId           string `json:"objId"`
}
