package resp

import "core/types"

type CategoryItemResp struct {
	CategoryId     int64      `json:"categoryId"`
	AssetType      int8       `json:"assetType"`
	Title          string     `json:"title"`
	Cover          string     `json:"cover"`
	ObjId          int64      `json:"objId"`
	ChapterId      int64      `json:"chapterId"`
	UpgradeChapter string     `json:"upgradeChapter"`
	UpdatedAt      types.Time `json:"updatedAt"`
	IsUpgrade      int8       `json:"isUpgrade"`
}
