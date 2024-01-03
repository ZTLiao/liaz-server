package controller

import (
	"business/handler"
	"business/storage"
	"core/redis"
	"core/system"
	"core/web"
)

type ComicController struct {
}

var _ web.IWebController = &ComicController{}

func (e *ComicController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var comicHandler = &handler.ComicHandler{
		ComicDb:                storage.NewComicDb(db),
		ComicChapterDb:         storage.NewComicChapterDb(db),
		ComicChapterItemDb:     storage.NewComicChapterItemDb(db),
		ComicSubscribeDb:       storage.NewComicSubscribeDb(db),
		BrowseDb:               storage.NewBrowseDb(db),
		ComicSubscribeNumCache: storage.NewComicSubscribeNumCache(redis),
		ComicHitNumCache:       storage.NewComicHitNumCache(redis),
	}
	iWebRoutes.GET("/comic/:comicId", comicHandler.ComicDetail)
	iWebRoutes.GET("/comic/upgrade", comicHandler.ComicUpgrade)
	iWebRoutes.GET("/comic/catalogue", comicHandler.ComicCatalogue)
	iWebRoutes.GET("/comic/get", comicHandler.GetComic)
}
