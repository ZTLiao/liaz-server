package handler

import (
	basicHandler "basic/handler"
	"business/resp"
	"core/constant"
	"core/response"
	"core/web"
	"encoding/json"
)

type ClientHandler struct {
	SysConfHandler *basicHandler.SysConfHandler
}

func (e *ClientHandler) ClientInit(wc *web.WebContext) interface{} {
	fileUrl, err := e.SysConfHandler.GetConfValueByKey(constant.FILE_URL)
	if err != nil {
		wc.AbortWithError(err)
	}
	var key = new(resp.KeyConfig)
	var app = new(resp.AppConfig)
	app.FileUrl = fileUrl
	//格式化
	appJson, err := json.Marshal(app)
	if err != nil {
		return err
	}
	var clientInitResp = &resp.ClientInitResp{
		Key: key,
		App: string(appJson),
	}
	return response.ReturnOK(clientInitResp)
}
