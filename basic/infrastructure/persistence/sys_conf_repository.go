package persistence

import (
	"basic/domain/entity"
	"basic/domain/repository"
	"core/constant"
	"core/logger"

	"github.com/go-xorm/xorm"
)

type SysConfReop struct {
	db *xorm.Engine
}

var _ repository.SysConfRepository = &SysConfReop{}

func NewSysConfRepository(db *xorm.Engine) *SysConfReop {
	return &SysConfReop{db}
}

func (e *SysConfReop) SaveSysConf(sysConf *entity.SysConf) (*entity.SysConf, map[string]string) {
	dbErr := map[string]string{}
	_, err := e.db.Insert(&sysConf)
	if err != nil {
		dbErr[constant.DB_ERROR] = err.Error()
		logger.Error(err.Error())
		return nil, dbErr
	}
	return sysConf, nil
}
func (e *SysConfReop) GetSysConf(confId int64) (*entity.SysConf, error) {
	var sysConf entity.SysConf
	_, err := e.db.ID(confId).Get(&sysConf)
	if err != nil {
		return nil, err
	}
	return &sysConf, nil
}
func (e *SysConfReop) UpdateSysConf(sysConf *entity.SysConf) (*entity.SysConf, map[string]string) {
	dbErr := map[string]string{}
	_, err := e.db.ID(sysConf.ConfId).Update(&sysConf)
	if err != nil {
		dbErr[constant.DB_ERROR] = err.Error()
		logger.Error(err.Error())
		return nil, dbErr
	}
	return sysConf, nil
}
func (e *SysConfReop) DeleteSysConf(confId int64) error {
	var sysConf entity.SysConf
	_, err := e.db.ID(confId).Delete(&sysConf)
	if err != nil {
		return err
	}
	return nil
}
