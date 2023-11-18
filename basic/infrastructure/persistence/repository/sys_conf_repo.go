package repository

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"
	"basic/infrastructure/persistence/repository/memory"
	"basic/infrastructure/persistence/repository/store"
)

type SysConfReop struct {
	sysConfStore  *store.SysConfStore
	sysConfMemory *memory.SysConfMemory
}

var _ repository.SysConfRepoInterface = &SysConfReop{}

func NewSysConfRepo(sysConfStore *store.SysConfStore, sysConfMemory *memory.SysConfMemory) *SysConfReop {
	return &SysConfReop{sysConfStore, sysConfMemory}
}

func (e *SysConfReop) SaveSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	sysConf, err := e.sysConfStore.SaveSysConf(sysConf)
	if err != nil {
		return nil, err
	}
	err = e.sysConfMemory.HDel(sysConf.ConfKey)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfReop) GetSysConf(confId int64) (*entity.SysConf, error) {
	return e.sysConfStore.GetSysConf(confId)
}

func (e *SysConfReop) UpdateSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	sysConf, err := e.sysConfStore.UpdateSysConf(sysConf)
	if err != nil {
		return nil, err
	}
	err = e.sysConfMemory.HDel(sysConf.ConfKey)
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfReop) DeleteSysConf(confId int64) error {
	sysConf, err := e.GetSysConf(confId)
	if err != nil {
		return err
	}
	err = e.sysConfStore.DeleteSysConf(confId)
	if err != nil {
		return err
	}
	err = e.sysConfMemory.HDel(sysConf.ConfKey)
	if err != nil {
		return err
	}
	return nil
}

func (e *SysConfReop) GetSysConfList() ([]entity.SysConf, error) {
	return e.sysConfStore.GetSysConfList()
}

func (e *SysConfReop) GetSysConfByKey(confKey string) (*entity.SysConf, error) {
	sysConf, err := e.sysConfMemory.HGet(confKey)
	if err != nil {
		return nil, err
	}
	if sysConf == nil {
		sysConf, err = e.sysConfStore.GetSysConfByKey(confKey)
		if err != nil {
			return nil, err
		}
		if sysConf != nil {
			e.sysConfMemory.HSet(confKey, sysConf)
		}
	}
	return sysConf, nil
}
