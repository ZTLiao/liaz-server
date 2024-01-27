package model

import (
	"core/model"
	"core/types"
)

type Notification struct {
	NotificationId int64      `json:"notificationId" xorm:"notification_id pk autoincr BIGINT"`
	Title          string     `json:"title" xorm:"title"`
	Content        string     `json:"content" xorm:"content"`
	Status         int8       `json:"status" xorm:"status"`
	CreatedAt      types.Time `json:"createdAt" xorm:"created_at timestamp created"`
	UpdatedAt      types.Time `json:"updatedAt" xorm:"updated_at timestamp updated"`
}

var _ model.BaseModel = &Notification{}

func (e *Notification) TableName() string {
	return "notification"
}
