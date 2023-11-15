package repository

import "basic/domain/entity"

type SysConfRepository interface {
	SaveSysConf(*entity.SysConf) (*entity.SysConf, map[string]string)
	GetSysConf(int64) (*entity.SysConf, error)
	UpdateSysConf(*entity.SysConf) (*entity.SysConf, map[string]string)
	DeleteSysConf(int64) error
}
