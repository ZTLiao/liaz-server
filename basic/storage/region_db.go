package storage

import "github.com/go-xorm/xorm"

type RegionDb struct {
	db *xorm.Engine
}

func NewRegionDb(db *xorm.Engine) *RegionDb {
	return &RegionDb{db}
}
