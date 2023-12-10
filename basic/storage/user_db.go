package storage

import (
	"basic/enums"
	"basic/model"
	"core/utils"

	"github.com/go-xorm/xorm"
)

type UserDb struct {
	db *xorm.Engine
}

func NewUserDb(db *xorm.Engine) *UserDb {
	return &UserDb{db}
}

func (e *UserDb) UpdateLocation(userId int64, ipAddr string) error {
	if len(ipAddr) == 0 {
		return nil
	}
	country, province, city, err := utils.GetLocation(ipAddr)
	if err != nil {
		return err
	}
	if len(country) == 0 && len(province) == 0 && len(city) == 0 {
		return nil
	}
	var user = new(model.User)
	user.UserId = userId
	user.Country = country
	user.Province = province
	user.City = city
	_, err = e.db.ID(userId).Cols("country", "province", "city").Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (e *UserDb) SignUp(userId int64, nickname string, avatar string, gender int8, userType int8) error {
	var user = new(model.User)
	user.UserId = userId
	user.Nickname = nickname
	user.Avatar = avatar
	user.Gender = gender
	user.Type = userType
	user.Status = enums.USER_STATUS_OF_ENABLE
	count, err := e.db.Where("user_id = ?", userId).Count(user)
	if err != nil {
		return err
	}
	if count == 0 {
		_, err = e.db.Insert(user)
		if err != nil {
			return err
		}
	} else {
		_, err = e.db.ID(userId).Update(user)
		if err != nil {
			return err
		}
	}
	return nil
}
