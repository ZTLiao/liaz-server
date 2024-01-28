package storage

import (
	"basic/model"
	"core/constant"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type SysConfDb struct {
	db *xorm.Engine
}

func NewSysConfDb(db *xorm.Engine) *SysConfDb {
	return &SysConfDb{db}
}

func (e *SysConfDb) SaveSysConf(sysConf *model.SysConf) (*model.SysConf, error) {
	var now = types.Time(time.Now())
	sysConf.CreatedAt = now
	sysConf.UpdatedAt = now
	_, err := e.db.Insert(sysConf)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfDb) GetSysConf(confId int64) (*model.SysConf, error) {
	var sysConf model.SysConf
	_, err := e.db.ID(confId).Get(&sysConf)
	if err != nil {
		return nil, err
	}
	if sysConf.ConfId == 0 {
		return nil, nil
	}
	return &sysConf, nil
}

func (e *SysConfDb) UpdateSysConf(sysConf *model.SysConf) (*model.SysConf, error) {
	_, err := e.db.ID(sysConf.ConfId).Update(sysConf)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfDb) DeleteSysConf(confId int64) error {
	var sysConf model.SysConf
	_, err := e.db.ID(confId).Delete(&sysConf)
	if err != nil {
		return err
	}
	return nil
}

func (e *SysConfDb) GetSysConfList() ([]model.SysConf, error) {
	var sysConfs []model.SysConf
	err := e.db.OrderBy("created_at asc").Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	return sysConfs, nil
}

func (e *SysConfDb) GetSysConfByKey(confKey string) (*model.SysConf, error) {
	var sysConfs []model.SysConf
	err := e.db.Where("conf_key = ? and status = ?", confKey, constant.YES).Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	if len(sysConfs) > 0 {
		return &sysConfs[0], nil
	}
	return nil, nil
}

func (e *SysConfDb) GetSysConfByType(confType int8) ([]model.SysConf, error) {
	var sysConfs []model.SysConf
	err := e.db.SQL(
		`
		select 
			sc.conf_id,
			sc.conf_key,
			sc.conf_name,
			sc.conf_kind,
			sc.conf_type,
			sc.conf_value,
			sc.description,
			sc.status,
			sc.created_at,
			sc.updated_at
		from 
			sys_conf as sc
		where 
			sc.status = 1
			and (sc.conf_type & ?) != 0
		`, confType).Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	return sysConfs, nil
}
