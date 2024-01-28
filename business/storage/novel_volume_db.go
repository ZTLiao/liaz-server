package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type NovelVolumeDb struct {
	db *xorm.Engine
}

func NewNovelVolumeDb(db *xorm.Engine) *NovelVolumeDb {
	return &NovelVolumeDb{db}
}

func (e *NovelVolumeDb) GetNovelVolumeById(novelVolumeId int64) (*model.NovelVolume, error) {
	var novelVolume model.NovelVolume
	_, err := e.db.ID(novelVolumeId).Get(&novelVolume)
	if err != nil {
		return nil, err
	}
	if novelVolume.NovelVolumeId == 0 {
		return nil, nil
	}
	return &novelVolume, nil
}
