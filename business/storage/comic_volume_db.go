package storage

import (
	"business/model"

	"github.com/go-xorm/xorm"
)

type ComicVolumeDb struct {
	db *xorm.Engine
}

func NewComicVolumeDb(db *xorm.Engine) *ComicVolumeDb {
	return &ComicVolumeDb{db}
}

func (e *ComicVolumeDb) GetNovelVolumeById(comicVolumeId int64) (*model.ComicVolume, error) {
	var comicVolume model.ComicVolume
	_, err := e.db.ID(comicVolumeId).Get(&comicVolume)
	if err != nil {
		return nil, err
	}
	if comicVolume.ComicVolumeId == 0 {
		return nil, nil
	}
	return &comicVolume, nil
}
