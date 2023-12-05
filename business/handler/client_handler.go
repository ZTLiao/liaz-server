package handler

import (
	basicHandler "basic/handler"
	"business/resp"
	"core/config"
	"core/constant"
	"core/response"
	"core/utils"
	"core/web"
	"encoding/json"
)

type ClientHandler struct {
	SysConfHandler *basicHandler.SysConfHandler
	SecurityConfig *config.Security
}

func (e *ClientHandler) ClientInit(wc *web.WebContext) interface{} {
	fileUrl, err := e.SysConfHandler.GetConfValueByKey(constant.FILE_URL)
	if err != nil {
		wc.AbortWithError(err)
	}
	//获取加密密钥
	var key = new(resp.KeyConfig)
	//接口加签
	signKey := e.SecurityConfig.SignKey
	if len(signKey) > 0 {
		signKey = utils.EncryptKey(signKey)
	}
	key.K1 = signKey
	//数据加密
	key.K2 = e.SecurityConfig.PublicKey
	var app = new(resp.AppConfig)
	app.FileUrl = fileUrl
	//格式化
	appJson, err := json.Marshal(app)
	if err != nil {
		wc.AbortWithError(err)
	}
	//加密
	if e.SecurityConfig.Encrypt {
		encryptPlain, err := utils.PriKeyEncrypt(string(appJson), e.SecurityConfig.PrivateKey)
		if err != nil {
			wc.AbortWithError(err)
		}
		appJson = []byte(encryptPlain)
	}
	var clientInitResp = &resp.ClientInitResp{
		Key: key,
		App: string(appJson),
	}
	return response.ReturnOK(clientInitResp)
}
