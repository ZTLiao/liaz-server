package model

import "core/types"

type User struct {
	UserId      int64      `json:"userId" xorm:"user_id pk BIGINT"`
	Nickname    string     `json:"nickname" xorm:"nickname"`
	Phone       string     `json:"phone" xorm:"phone"`
	Email       string     `json:"email" xorm:"email"`
	Avatar      string     `json:"avatar" xorm:"avatar"`
	Description string     `json:"description" xorm:"description"`
	Gender      int8       `json:"gender" xorm:"gender"`
	UserType    int8       `json:"userType" xorm:"user_type"`
	Status      int8       `json:"status" xorm:"status"`
	CreatedAt   types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt   types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *User) TableName() string {
	return "user"
}
