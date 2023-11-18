package repository

import "basic/infrastructure/persistence/entity"

type SysConfRepoInterface interface {
	SaveSysConf(*entity.SysConf) (*entity.SysConf, error)
	GetSysConf(int64) (*entity.SysConf, error)
	UpdateSysConf(*entity.SysConf) (*entity.SysConf, error)
	DeleteSysConf(int64) error
	GetSysConfList() ([]entity.SysConf, error)
	GetSysConfByKey(string) (*entity.SysConf, error)
}
