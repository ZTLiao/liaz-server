package store

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"
	"core/constant"

	"github.com/go-xorm/xorm"
)

type SysConfStore struct {
	db *xorm.Engine
}

var _ repository.SysConfRepoInterface = &SysConfStore{}

func NewSysConfStore(db *xorm.Engine) *SysConfStore {
	return &SysConfStore{db}
}

func (e *SysConfStore) SaveSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	_, err := e.db.Insert(sysConf)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfStore) GetSysConf(confId int64) (*entity.SysConf, error) {
	var sysConf entity.SysConf
	_, err := e.db.ID(confId).Get(&sysConf)
	if err != nil {
		return nil, err
	}
	return &sysConf, nil
}

func (e *SysConfStore) UpdateSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	_, err := e.db.ID(sysConf.ConfId).Update(sysConf)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfStore) DeleteSysConf(confId int64) error {
	var sysConf entity.SysConf
	_, err := e.db.ID(confId).Delete(&sysConf)
	if err != nil {
		return err
	}
	return nil
}

func (e *SysConfStore) GetSysConfList() ([]entity.SysConf, error) {
	var sysConfs []entity.SysConf
	err := e.db.OrderBy("created_at asc").Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	return sysConfs, nil
}

func (e *SysConfStore) GetSysConfByKey(confKey string) (*entity.SysConf, error) {
	var sysConfs []entity.SysConf
	err := e.db.Where("conf_key = ? and status = ?", confKey, constant.YES).Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	if len(sysConfs) > 0 {
		return &sysConfs[0], nil
	}
	return nil, nil
}

func (e *SysConfStore) GetSysConfByKind(confKind int8) ([]entity.SysConf, error) {
	var sysConfs []entity.SysConf
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
			and (sc.conf_kind & ?) != 0
		`, confKind).Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	return sysConfs, nil
}
