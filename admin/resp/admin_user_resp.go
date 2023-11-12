package resp

import (
	"admin/model"
	"core/types"
)

type AdminUserResp struct {
	model.AdminUser
	LastTime types.Time `json:"lastTime"` //上次登录时间
}
