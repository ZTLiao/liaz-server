package controller

import (
	"admin/handler"
	"basic/storage"
	"core/system"
	"core/web"
)

type AdminAuthorController struct {
}

var _ web.IWebController = &AdminAuthorController{}

func (e *AdminAuthorController) Router(iWebRoutes web.IWebRoutes) {
	db := system.GetXormEngine()
	var adminAuthorHandler = handler.AdminAuthorHandler{
		AuthorDb: storage.NewAuthorDb(db),
	}
	iWebRoutes.GET("/author/page", adminAuthorHandler.GetAuthorPage)
	iWebRoutes.GET("/author", adminAuthorHandler.GetAuthorList)
	iWebRoutes.POST("/author", adminAuthorHandler.SaveAuthor)
	iWebRoutes.PUT("/author", adminAuthorHandler.UpdateAuthor)
	iWebRoutes.DELETE("/author/:authorId", adminAuthorHandler.DelAuthor)
}
