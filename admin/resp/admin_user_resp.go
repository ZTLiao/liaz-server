package resp

import (
	"core/types"
)

type AdminUserResp struct {
	AdminId  int64      `json:"adminId"`  //用户ID
	Name     string     `json:"name"`     //名称
	Username string     `json:"username"` //账号
	Avatar   string     `json:"avatar"`   //头像
	LastTime types.Time `json:"lastTime"` //上次登录时间
}
