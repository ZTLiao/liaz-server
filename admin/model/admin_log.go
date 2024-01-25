package model

import (
	"core/model"
	"core/types"
)

type AdminLog struct {
	LogId     int64      `json:"logId" xorm:"log_id pk autoincr BIGINT"`
	AdminId   int64      `json:"adminId" xorm:"admin_id"`
	Uri       string     `json:"uri" xorm:"uri"`
	Params    string     `json:"params" xorm:"params"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
}

var _ model.BaseModel = &AdminLog{}

func (e *AdminLog) TableName() string {
	return "admin_log"
}
