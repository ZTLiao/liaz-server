package resp

type CategoryItemResp struct {
	CategoryId     int64  `json:"categoryId"`
	AssetType      int8   `json:"assetType"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	ObjId          int64  `json:"objId"`
	UpgradeChapter string `json:"upgradeChapter"`
}
