package model

import "core/types"

type Account struct {
	UserId    int64      `json:"userId" xorm:"user_id pk autoincr BIGINT"`
	Username  string     `json:"username" xorm:"username"`
	Phone     string     `json:"phone" xorm:"phone"`
	Email     string     `json:"email" xorm:"email"`
	Password  string     `json:"password" xorm:"password"`
	Flag      int8       `json:"flag" xorm:"flag"`
	Status    int8       `json:"status" xorm:"status"`
	CreatedAt types.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt types.Time `json:"updatedAt" xorm:"updated_at"`
}

func (e *Account) TableName() string {
	return "account"
}
