package transfer

type SearchDto struct {
	AssetId        int64  `xorm:"asset_id"`
	ObjId          int64  `xorm:"obj_id"`
	Title          string `xorm:"title"`
	Cover          string `xorm:"cover"`
	AssetType      int8   `xorm:"asset_type"`
	Categories     string `xorm:"categories"`
	Authors        string `xorm:"authors"`
	UpgradeChapter string `xorm:"upgrade_chapter"`
}
