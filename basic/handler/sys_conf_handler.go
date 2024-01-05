package handler

import (
	"basic/model"
	"basic/storage"
	"core/constant"
	"core/logger"
	"strconv"
)

type SysConfHandler struct {
	sysConfDb    *storage.SysConfDb
	sysConfCache *storage.SysConfCache
}

func NewSysConfHandler(sysConfDb *storage.SysConfDb, sysConfCache *storage.SysConfCache) *SysConfHandler {
	return &SysConfHandler{sysConfDb, sysConfCache}
}

func (e *SysConfHandler) SaveSysConf(sysConf *model.SysConf) (*model.SysConf, error) {
	sysConf, err := e.sysConfDb.SaveSysConf(sysConf)
	if err != nil {
		return nil, err
	}
	err = e.sysConfCache.Del()
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfHandler) GetSysConf(confId int64) (*model.SysConf, error) {
	return e.sysConfDb.GetSysConf(confId)
}

func (e *SysConfHandler) UpdateSysConf(sysConf *model.SysConf) (*model.SysConf, error) {
	sysConf, err := e.sysConfDb.UpdateSysConf(sysConf)
	if err != nil {
		return nil, err
	}
	err = e.sysConfCache.Del()
	if err != nil {
		return nil, err
	}
	return sysConf, nil
}

func (e *SysConfHandler) DeleteSysConf(confId int64) error {
	err := e.sysConfDb.DeleteSysConf(confId)
	if err != nil {
		return err
	}
	err = e.sysConfCache.Del()
	if err != nil {
		return err
	}
	return nil
}

func (e *SysConfHandler) GetSysConfList() ([]model.SysConf, error) {
	return e.sysConfDb.GetSysConfList()
}

func (e *SysConfHandler) GetSysConfByKey(confKey string) (*model.SysConf, error) {
	sysConf, err := e.sysConfCache.HGet(confKey)
	if err != nil {
		return nil, err
	}
	if sysConf == nil {
		sysConf, err = e.sysConfDb.GetSysConfByKey(confKey)
		if err != nil {
			return nil, err
		}
		sysConfs, err := e.GetSysConfList()
		if err != nil {
			return nil, err
		}
		for _, v := range sysConfs {
			e.sysConfCache.HSet(v.ConfKey, &v)
		}
	}
	return sysConf, nil
}

func (e *SysConfHandler) GetSysConfByType(confType int8) ([]model.SysConf, error) {
	sysConfMap, err := e.sysConfCache.HGetAll()
	if err != nil {
		return nil, err
	}
	if len(sysConfMap) == 0 {
		sysConfs, err := e.sysConfDb.GetSysConfByType(confType)
		if err != nil {
			return nil, err
		}
		for _, v := range sysConfs {
			if len(v.ConfKey) > 0 {
				e.sysConfCache.HSet(v.ConfKey, &v)

			}
		}
		sysConfMap, err = e.sysConfCache.HGetAll()
		if err != nil {
			return nil, err
		}
	}
	var sysConfs = make([]model.SysConf, 0)
	if len(sysConfMap) > 0 {
		for _, v := range sysConfMap {
			if (v.ConfType&confType) != 0 && v.Status == constant.YES && len(v.ConfKey) > 0 {
				sysConfs = append(sysConfs, v)
			}
		}
	}
	return sysConfs, nil
}

func (e *SysConfHandler) GetConfValueByKey(confKey string) (string, error) {
	sysConf, err := e.GetSysConfByKey(confKey)
	if err != nil {
		logger.Error(err.Error())
		return "", nil
	}
	if sysConf == nil {
		return "", nil
	}
	return sysConf.ConfValue, nil
}

func (e *SysConfHandler) GetIntValueByKey(confKey string) (int, error) {
	sysConf, err := e.GetSysConfByKey(confKey)
	if err != nil {
		logger.Error(err.Error())
		return 0, nil
	}
	if sysConf == nil {
		return 0, nil
	}
	confValue := sysConf.ConfValue
	if len(confValue) == 0 {
		return 0, nil
	}
	intValue, err := strconv.ParseInt(confValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(intValue), nil
}
