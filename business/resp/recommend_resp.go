package resp

type RecommendResp struct {
	RecommendId int64               `json:"recommendId"`
	Title       string              `json:"title"`
	ShowType    int8                `json:"showType"`
	IsShowTitle bool                `json:"isShowTitle"`
	OptType     int8                `json:"optType"`
	OptValue    string              `json:"optValue"`
	SeqNo       int                 `json:"seqNo"`
	Items       []RecommendItemResp `json:"items"`
}
