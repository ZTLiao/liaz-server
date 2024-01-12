package controller

import (
	"business/handler"
	"business/listener"
	"business/storage"
	"core/constant"
	"core/event"
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
	var comicDb = storage.NewComicDb(db)
	var comicSubscribeDb = storage.NewComicSubscribeDb(db)
	var comicHitNumCache = storage.NewComicHitNumCache(redis)
	event.Bus.Subscribe(constant.COMIC_HIT_TOPIC, listener.NewComicHitListener(comicDb, comicSubscribeDb, comicHitNumCache))
	event.Bus.Subscribe(constant.COMIC_POPULAR_RANK_TOPIC, listener.NewComicPopularRankListener(storage.NewComicRankCache(redis)))
	var comicHandler = &handler.ComicHandler{
		ComicDb:                comicDb,
		ComicChapterDb:         storage.NewComicChapterDb(db),
		ComicChapterItemDb:     storage.NewComicChapterItemDb(db),
		ComicSubscribeDb:       comicSubscribeDb,
		BrowseDb:               storage.NewBrowseDb(db),
		ComicSubscribeNumCache: storage.NewComicSubscribeNumCache(redis),
		ComicHitNumCache:       comicHitNumCache,
		ComicDetailCache:       storage.NewComicDetailCache(redis),
		ComicUpgradeItemCache:  storage.NewComicUpgradeItemCache(redis),
	}
	iWebRoutes.GET("/comic/:comicId", comicHandler.ComicDetail)
	iWebRoutes.GET("/comic/upgrade", comicHandler.ComicUpgrade)
	iWebRoutes.GET("/comic/catalogue", comicHandler.ComicCatalogue)
	iWebRoutes.GET("/comic/get", comicHandler.GetComic)
}
