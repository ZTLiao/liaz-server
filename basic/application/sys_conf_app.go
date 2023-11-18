package application

import (
	"basic/domain/repository"
	"basic/infrastructure/persistence/entity"
)

type sysConfApp struct {
	sysConfRepo repository.SysConfRepoInterface
}

var _ SysConfAppInterface = &sysConfApp{}

type SysConfAppInterface interface {
	SaveSysConf(*entity.SysConf) (*entity.SysConf, error)
	GetSysConf(int64) (*entity.SysConf, error)
	UpdateSysConf(*entity.SysConf) (*entity.SysConf, error)
	DeleteSysConf(int64) error
	GetSysConfList() ([]entity.SysConf, error)
	GetSysConfByKey(confKey string) (*entity.SysConf, error)
}

func (e *sysConfApp) SaveSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	return e.sysConfRepo.SaveSysConf(sysConf)
}

func (e *sysConfApp) GetSysConf(confId int64) (*entity.SysConf, error) {
	return e.sysConfRepo.GetSysConf(confId)
}

func (e *sysConfApp) UpdateSysConf(sysConf *entity.SysConf) (*entity.SysConf, error) {
	return e.sysConfRepo.UpdateSysConf(sysConf)
}

func (e *sysConfApp) DeleteSysConf(confId int64) error {
	return e.sysConfRepo.DeleteSysConf(confId)
}

func (e *sysConfApp) GetSysConfList() ([]entity.SysConf, error) {
	return e.sysConfRepo.GetSysConfList()
}

func (e *sysConfApp) GetSysConfByKey(confKey string) (*entity.SysConf, error) {
	return e.sysConfRepo.GetSysConfByKey(confKey)
}
