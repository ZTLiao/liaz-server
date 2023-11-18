package store

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"

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
	_, err := e.db.Insert(&sysConf)
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
	_, err := e.db.ID(sysConf.ConfId).Update(&sysConf)
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
	err := e.db.Where("conf_key = ?", confKey).Find(&sysConfs)
	if err != nil {
		return nil, err
	}
	if len(sysConfs) > 0 {
		return &sysConfs[0], nil
	}
	return nil, nil
}
