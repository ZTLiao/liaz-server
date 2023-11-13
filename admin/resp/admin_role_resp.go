package resp

type AdminRoleResp struct {
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	Checked  bool   `json:"checked"`
}
