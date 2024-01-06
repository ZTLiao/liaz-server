package transfer

type SearchDto struct {
	ObjId      int64  `xorm:"obj_id"`
	Title      string `xorm:"title"`
	Cover      string `xorm:"cover"`
	AssetType  int8   `xorm:"asset_type"`
	Categories string `xorm:"categories"`
	Authors    string `xorm:"authors"`
}
