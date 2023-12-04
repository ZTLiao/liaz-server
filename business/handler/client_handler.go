package handler

import (
	basic "basic/handler"
	"business/resp"
	"core/constant"
	"core/response"
	"core/web"
)

type ClientHandler struct {
	SysConfHandler *basic.SysConfHandler
}

func (e *ClientHandler) ClientInit(wc *web.WebContext) interface{} {
	fileUrl, err := e.SysConfHandler.GetConfValueByKey(constant.FILE_URL)
	if err != nil {
		wc.AbortWithError(err)
	}
	var clientInitResp = &resp.ClientInitResp{
		FileUrl: fileUrl,
	}
	return response.ReturnOK(clientInitResp)
}
