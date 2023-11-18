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
	err = e.sysConfMemory.Del()
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
	err = e.sysConfMemory.Del()
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfReop) DeleteSysConf(confId int64) error {
	err := e.sysConfStore.DeleteSysConf(confId)
	if err != nil {
		return err
	}
	err = e.sysConfMemory.Del()
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
		sysConfs, err := e.GetSysConfList()
		if err != nil {
			return nil, err
		}
		for _, v := range sysConfs {
			e.sysConfMemory.HSet(v.ConfKey, &v)
		}
	}
	return sysConf, nil
}

func (e *SysConfReop) GetSysConfByKind(confKind int8) ([]entity.SysConf, error) {
	sysConfMap, err := e.sysConfMemory.HGetAll()
	if err != nil {
		return nil, err
	}
	if len(sysConfMap) == 0 {
		sysConfs, err := e.sysConfStore.GetSysConfByKind(confKind)
		if err != nil {
			return nil, err
		}
		for _, v := range sysConfs {
			if len(v.ConfKey) > 0 {
				e.sysConfMemory.HSet(v.ConfKey, &v)

			}
		}
		sysConfMap, err = e.sysConfMemory.HGetAll()
		if err != nil {
			return nil, err
		}
	}
	var sysConfs = make([]entity.SysConf, 0)
	if len(sysConfMap) > 0 {
		for _, v := range sysConfMap {
			if (v.ConfKind&confKind) != 0 && len(v.ConfKey) > 0 {
				sysConfs = append(sysConfs, v)
			}
		}
	}
	return sysConfs, nil
}
