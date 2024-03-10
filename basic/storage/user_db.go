package storage

import (
	"basic/enums"
	"basic/model"
	"core/types"
	"core/utils"
	"time"

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
	var now = types.Time(time.Now())
	var user = new(model.User)
	user.UserId = userId
	user.Country = country
	user.Province = province
	user.City = city
	user.UpdatedAt = now
	_, err = e.db.ID(userId).Cols("country", "province", "city", "updatedAt").Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (e *UserDb) SignUp(userId int64, nickname string, avatar string, gender int8, userType int8) error {
	var now = types.Time(time.Now())
	var user = new(model.User)
	user.UserId = userId
	user.Nickname = nickname
	user.Avatar = avatar
	user.Gender = gender
	user.Type = userType
	user.Status = enums.USER_STATUS_OF_ENABLE
	user.CreatedAt = now
	user.UpdatedAt = now
	count, err := e.db.Where("user_id = ?", userId).Count(&model.User{})
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

func (e *UserDb) GetUserById(userId int64) (*model.User, error) {
	var user model.User
	_, err := e.db.ID(userId).Get(&user)
	if err != nil {
		return nil, err
	}
	if user.UserId == 0 {
		return nil, nil
	}
	return &user, nil
}

func (e *UserDb) UpdateUser(userId int64, avatar string, nickname string, phone string, email string, gender int8, description string) error {
	var now = types.Time(time.Now())
	_, err := e.db.ID(userId).Cols("avatar", "nickname", "phone", "email", "gender", "description").Update(&model.User{
		Avatar:      avatar,
		Nickname:    nickname,
		Phone:       phone,
		Email:       email,
		Gender:      gender,
		Description: description,
		UpdatedAt:   now,
	})
	return err
}
