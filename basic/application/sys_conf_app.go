package application

import (
	"basic/domain/entity"
	"basic/domain/repository"
)

type sysConfApp struct {
	sysConfRepo repository.SysConfRepository
}

var _ SysConfAppInterface = &sysConfApp{}

type SysConfAppInterface interface {
	SaveSysConf(*entity.SysConf) (*entity.SysConf, map[string]string)
	GetSysConf(int64) (*entity.SysConf, error)
	UpdateSysConf(*entity.SysConf) (*entity.SysConf, map[string]string)
	DeleteSysConf(int64) error
}

func (e *sysConfApp) SaveSysConf(sysConf *entity.SysConf) (*entity.SysConf, map[string]string) {
	return e.sysConfRepo.SaveSysConf(sysConf)
}
func (e *sysConfApp) GetSysConf(confId int64) (*entity.SysConf, error) {
	return e.sysConfRepo.GetSysConf(confId)
}
func (e *sysConfApp) UpdateSysConf(sysConf *entity.SysConf) (*entity.SysConf, map[string]string) {
	return e.sysConfRepo.UpdateSysConf(sysConf)
}
func (e *sysConfApp) DeleteSysConf(confId int64) error {
	return e.sysConfRepo.DeleteSysConf(confId)
}
