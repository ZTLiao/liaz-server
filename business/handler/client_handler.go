package handler

import (
	"basic/device"
	basicHandler "basic/handler"
	"business/resp"
	"core/config"
	"core/constant"
	"core/event"
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
	//获取加密密钥
	var key = new(resp.KeyConfig)
	//接口加签
	signKey := e.SecurityConfig.SignKey
	if len(signKey) > 0 {
		signKey = utils.EncryptKey(signKey)
	}
	publicKey := e.SecurityConfig.PublicKey
	if len(publicKey) > 0 {
		publicKey = utils.EncryptKey(publicKey)
	}
	key.K1 = signKey
	//数据加密
	key.K2 = publicKey
	//获取配置
	appJson := e.buildAppConfig(wc)
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
	//APP初始化事件
	event.Bus.Publish(constant.CLIENT_INIT_TOPIC, device.GetDeviceInfo(wc))
	return response.ReturnOK(clientInitResp)
}

func (e *ClientHandler) buildAppConfig(wc *web.WebContext) []byte {
	fileUrl, err := e.SysConfHandler.GetConfValueByKey(constant.FILE_URL)
	if err != nil {
		wc.AbortWithError(err)
	}
	var app = new(resp.AppConfig)
	app.FileUrl = fileUrl
	//格式化
	appJson, err := json.Marshal(app)
	if err != nil {
		wc.AbortWithError(err)
	}
	return appJson
}
