package storage

import (
	"business/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type DiscussDb struct {
	db *xorm.Engine
}

func NewDiscussDb(db *xorm.Engine) *DiscussDb {
	return &DiscussDb{db}
}

func (e *DiscussDb) Save(parentId int64, userId int64, objId int64, objType int8, content string, status int8) (int64, error) {
	var now = types.Time(time.Now())
	var discuss = &model.Discuss{
		ParentId:  parentId,
		UserId:    userId,
		ObjId:     objId,
		ObjType:   objType,
		Content:   content,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err := e.db.Insert(discuss)
	if err != nil {
		return 0, err
	}
	return discuss.DiscussId, nil
}

func (e *DiscussDb) GetDiscussPage(objId int64, objType int8, pageNum int32, pageSize int32) ([]model.Discuss, error) {
	var discusses []model.Discuss
	err := e.db.Where("obj_id = ? and obj_type = ? and status = ?", objId, objType, constant.YES).OrderBy("created_at desc").Limit(int(pageSize), int((pageNum-1)*pageSize)).Find(&discusses)
	if err != nil {
		return nil, err
	}
	return discusses, nil
}

func (e *DiscussDb) GetDiscussById(discussId int64) (*model.Discuss, error) {
	var discuss model.Discuss
	_, err := e.db.ID(discussId).Get(&discuss)
	if err != nil {
		return nil, err
	}
	return &discuss, nil
}
