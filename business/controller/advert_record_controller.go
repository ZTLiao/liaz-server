package controller

import (
	"business/handler"
	"business/storage"
	"core/system"
	"core/web"
)

type AdvertRecordController struct {
}

var _ web.IWebController = &AdvertRecordController{}

func (e *AdvertRecordController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var advertRecordHandler = handler.AdvertRecordHandler{
		AdvertRecordDb: storage.NewAdvertRecordDb(db),
	}
	iWebRoutes.POST("/advert/record", advertRecordHandler.Record)
}
