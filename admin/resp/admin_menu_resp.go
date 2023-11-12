package resp

type AdminMenuResp struct {
	MenuId   int64           `json:"menuId"`
	MenuName string          `json:"menuName"`
	ParentId int64           `json:"parentId"`
	Checked  bool            `json:"checked"`
	Childs   []AdminMenuResp `json:"childs,omitempty"`
}
