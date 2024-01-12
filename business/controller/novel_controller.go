package controller

import (
	basicStorage "basic/storage"
	"business/handler"
	"business/listener"
	businessStorage "business/storage"
	"core/constant"
	"core/event"
	"core/redis"
	"core/system"
	"core/web"
)

type NovelController struct {
}

var _ web.IWebController = &NovelController{}

func (e *NovelController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var redis = redis.NewRedisUtil(system.GetRedisClient())
	var novelDb = businessStorage.NewNovelDb(db)
	var novelSubscribeDb = businessStorage.NewNovelSubscribeDb(db)
	var novelHitNumCache = businessStorage.NewNovelHitNumCache(redis)
	event.Bus.Subscribe(constant.NOVEL_HIT_TOPIC, listener.NewNovelHitListener(novelDb, novelSubscribeDb, novelHitNumCache))
	event.Bus.Subscribe(constant.NOVEL_POPULAR_RANK_TOPIC, listener.NewNovelPopularRankListener(businessStorage.NewNovelRankCache(redis)))
	var novelHandler = handler.NovelHandler{
		NovelDb:                novelDb,
		NovelVolumeDb:          businessStorage.NewNovelVolumeDb(db),
		NovelChapterDb:         businessStorage.NewNovelChapterDb(db),
		NovelChapterItemDb:     businessStorage.NewNovelChapterItemDb(db),
		FileItemDb:             basicStorage.NewFileItemDb(db),
		NovelSubscribeDb:       novelSubscribeDb,
		BrowseDb:               businessStorage.NewBrowseDb(db),
		NovelSubscribeNumCache: businessStorage.NewNovelSubscribeNumCache(redis),
		NovelHitNumCache:       novelHitNumCache,
		NovelRankCache:         businessStorage.NewNovelRankCache(redis),
		NovelDetailCache:       businessStorage.NewNovelDetailCache(redis),
		NovelUpgradeItemCache:  businessStorage.NewNovelUpgradeItemCache(redis),
	}
	iWebRoutes.GET("/novel/:novelId", novelHandler.NovelDetail)
	iWebRoutes.GET("/novel/upgrade", novelHandler.NovelUpgrade)
	iWebRoutes.GET("/novel/catalogue", novelHandler.NovelCatalogue)
	iWebRoutes.GET("/novel/get", novelHandler.GetNovel)
}
