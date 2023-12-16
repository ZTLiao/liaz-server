package model

import (
	"core/model"
	"core/types"
)

type Author struct {
	AuthorId   int64      `json:"authorId" xorm:"author_id"`
	AuthorName string     `json:"authorName" xorm:"author_name"`
	SeqNo      int32      `json:"seqNo" xorm:"seq_no"`
	Status     int8       `json:"status" xorm:"status"`
	CreatedAt  types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt  types.Time `json:"updatedAt" xorm:"updated_at"`
}

var _ model.BaseModel = &Author{}

func (e *Author) TableName() string {
	return "author"
}
