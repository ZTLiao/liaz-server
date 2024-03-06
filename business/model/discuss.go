package model

import (
	"core/model"
	"core/types"
)

type Discuss struct {
	DiscussId int64      `json:"discussId" xorm:"discuss_id pk autoincr BIGINT"`
	ParentId  int64      `json:"parentId" xorm:"parent_id"`
	UserId    int64      `json:"userId" xorm:"user_id"`
	ObjId     int64      `json:"objId" xorm:"obj_id"`
	ObjType   int8       `json:"objType" xorm:"obj_type"`
	Content   string     `json:"content" xorm:"content"`
	Status    int8       `json:"status" xorm:"status"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &Discuss{}

func (e *Discuss) TableName() string {
	return "discuss"
}
