package storage

import (
	"basic/device"
	"business/model"
	"core/types"
	"core/utils"
	"time"

	"github.com/go-xorm/xorm"
)

type MessageBoardDb struct {
	db *xorm.Engine
}

func NewMessageBoardDb(db *xorm.Engine) *MessageBoardDb {
	return &MessageBoardDb{db}
}

func (e *MessageBoardDb) InsertMessageBoard(content string, deviceInfo *device.DeviceInfo) error {
	var now = types.Time(time.Now())
	clientIp := deviceInfo.ClientIp
	ipRegion, _ := utils.GetAddress(clientIp)
	var record = &model.MessageBoard{
		Content:    content,
		DeviceId:   deviceInfo.DeviceId,
		Os:         deviceInfo.Os,
		OsVersion:  deviceInfo.OsVersion,
		App:        deviceInfo.App,
		AppVersion: deviceInfo.AppVersion,
		Model:      deviceInfo.Model,
		Imei:       deviceInfo.Imei,
		Channel:    deviceInfo.Channel,
		IspType:    deviceInfo.IspType,
		NetType:    deviceInfo.NetType,
		ClientIp:   deviceInfo.ClientIp,
		IpRegion:   ipRegion,
		CreatedAt:  now,
	}
	_, err := e.db.Insert(record)
	if err != nil {
		return err
	}
	return nil
}

func (e *MessageBoardDb) GetMessageBoardPage(startRow int, endRow int) ([]model.MessageBoard, int64, error) {
	var messageBoards []model.MessageBoard
	err := e.db.OrderBy("created_at asc").Limit(endRow, startRow).Find(&messageBoards)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.MessageBoard{})
	if err != nil {
		return nil, 0, err
	}
	return messageBoards, total, nil
}
