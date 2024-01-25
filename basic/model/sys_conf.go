package model

import (
	"core/model"
	"core/types"
)

type SysConf struct {
	ConfId      int64      `json:"confId" xorm:"conf_id pk autoincr BIGINT"`
	ConfKey     string     `json:"confKey" xorm:"conf_key"`
	ConfName    string     `json:"confName" xorm:"conf_name"`
	ConfKind    int8       `json:"confKind" xorm:"conf_kind"`
	ConfType    int8       `json:"confType" xorm:"conf_type"`
	ConfValue   string     `json:"confValue" xorm:"conf_value"`
	Description string     `json:"description" xorm:"description"`
	Status      int8       `json:"status" xorm:"status"`
	CreatedAt   types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt   types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &SysConf{}

func (e *SysConf) TableName() string {
	return "sys_conf"
}
