package storage

import (
	"business/enums"
	"business/model"
	"core/types"
	"time"

	"github.com/go-xorm/xorm"
)

type AppVersionDb struct {
	db *xorm.Engine
}

func NewAppVersionDb(db *xorm.Engine) *AppVersionDb {
	return &AppVersionDb{db}
}

func (e *AppVersionDb) GetAppVersionPage(startRow int, endRow int) ([]model.AppVersion, int64, error) {
	var appVersions []model.AppVersion
	err := e.db.OrderBy("created_at desc").Limit(endRow, startRow).Find(&appVersions)
	if err != nil {
		return nil, 0, err
	}
	total, err := e.db.Count(&model.AppVersion{})
	if err != nil {
		return nil, 0, err
	}
	return appVersions, total, nil
}

func (e *AppVersionDb) SaveOrUpdateAppVersion(appVersion *model.AppVersion) error {
	var now = types.Time(time.Now())
	versionId := appVersion.VersionId
	if versionId == 0 {
		appVersion.CreatedAt = now
		_, err := e.db.Insert(appVersion)
		if err != nil {
			return err
		}
	} else {
		appVersion.UpdatedAt = now
		_, err := e.db.ID(versionId).Update(appVersion)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AppVersionDb) DelAppVersion(versionId int64) error {
	_, err := e.db.ID(versionId).Delete(&model.AppVersion{})
	if err != nil {
		return err
	}
	return nil
}

func (e *AppVersionDb) GetLatest(os string, channel string) (*model.AppVersion, error) {
	var appVersion model.AppVersion
	_, err := e.db.Where("os = ? and channel = ? and status in (?, ?, ?)", os, channel, enums.APP_VERSION_FOR_ONLINE, enums.APP_VERSION_FOR_SUGGEST, enums.APP_VERSION_FOR_FORCE).OrderBy("created_at desc").Limit(1, 0).Get(&appVersion)
	if err != nil {
		return nil, err
	}
	if appVersion.VersionId == 0 {
		return nil, nil
	}
	return &appVersion, nil
}