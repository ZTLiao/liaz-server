package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	businessStorage "business/storage"
	"core/system"
	"core/web"
)

type DiscussController struct {
}

var _ web.IWebController = &DiscussController{}

func (e *DiscussController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var discussHandler = handler.DiscussHandler{
		DiscussDb:         businessStorage.NewDiscussDb(db),
		DiscussResourceDb: businessStorage.NewDiscussResourceDb(db),
		UserDb:            basicStorage.NewUserDb(db),
	}
	iWebRoutes.POST("/discuss", discussHandler.Discuss)
}
