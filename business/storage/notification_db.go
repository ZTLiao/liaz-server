package storage

import (
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type NotificationDb struct {
	db *xorm.Engine
}

func NewNotificationDb(db *xorm.Engine) *NotificationDb {
	return &NotificationDb{db}
}

func (e *NotificationDb) GetNotificationPage(startRow int, endRow int) ([]model.Notification, int64, error) {
	var Notifications []model.Notification
	err := e.db.OrderBy("created_at desc").Limit(endRow, startRow).Find(&Notifications)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.Notification{})
	if err != nil {
		return nil, 0, err
	}
	return Notifications, total, nil
}

func (e *NotificationDb) SaveOrUpdateNotification(Notification *model.Notification) error {
	var now = types.Time(time.Now())
	notificationId := Notification.NotificationId
	if notificationId == 0 {
		Notification.CreatedAt = now
		_, err := e.db.Insert(Notification)
		if err != nil {
			return err
		}
	} else {
		Notification.UpdatedAt = now
		_, err := e.db.ID(notificationId).Update(Notification)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *NotificationDb) DelNotification(versionId int64) error {
	_, err := e.db.ID(versionId).Delete(&model.Notification{})
	if err != nil {
		return err
	}
	return nil
}
