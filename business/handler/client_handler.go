package handler

import (
	"basic/device"
	"basic/enums"
	basicHandler "basic/handler"
	"business/resp"
	"core/config"
	"core/constant"
	"core/event"
	"core/response"
	"core/utils"
	"core/web"
	"encoding/json"
	"strconv"
)

var (
	client *resp.ClientInitResp
)

type ClientHandler struct {
	SysConfHandler *basicHandler.SysConfHandler
	SecurityConfig *config.Security
}

func (e *ClientHandler) ClientInit(wc *web.WebContext) interface{} {
	var err error
	var clientInitResp *resp.ClientInitResp
	if client == nil {
		clientInitResp, err = e.buildClientConfig()
		if err != nil {
			wc.AbortWithError(err)
		}
	} else {
		clientInitResp = client
		go e.buildClientConfig()
	}
	//发布事件
	event.Bus.Publish(constant.CLIENT_INIT_TOPIC, device.GetDeviceInfo(wc))
	return response.ReturnOK(clientInitResp)
}

func (e *ClientHandler) buildClientConfig() (*resp.ClientInitResp, error) {
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
	appJson, err := e.buildAppConfig()
	if err != nil {
		return nil, err
	}
	//加密
	if e.SecurityConfig.Encrypt {
		encryptPlain, err := utils.PriKeyEncrypt(string(appJson), e.SecurityConfig.PrivateKey)
		if err != nil {
			return nil, err
		}
		appJson = []byte(encryptPlain)
	}
	var clientInitResp = &resp.ClientInitResp{
		Key: key,
		App: string(appJson),
	}
	client = clientInitResp
	return clientInitResp, nil
}

func (e *ClientHandler) buildAppConfig() ([]byte, error) {
	sysConfs, err := e.SysConfHandler.GetSysConfByType(int8(enums.CONF_TYPE_OF_COMMON | enums.CONF_TYPE_OF_CLIENT))
	if err != nil {
		return nil, err
	}
	if len(sysConfs) == 0 {
		return nil, nil
	}
	var confMap = make(map[string]any)
	//0 自定义 1 文本 2 布尔 3 JSON
	for _, v := range sysConfs {
		confKey := v.ConfKey
		confKind := v.ConfKind
		confValue := v.ConfValue
		if confKind == enums.CONF_KIND_OF_BOOL {
			boolValue, err := strconv.ParseBool(confValue)
			if err != nil {
				return nil, err
			}
			confMap[confKey] = boolValue
		} else if confKind == enums.CONF_KIND_OF_JSON {
			var jsonValue map[string]interface{}
			err := json.Unmarshal([]byte(confValue), &jsonValue)
			if err != nil {
				return nil, err
			}
			confMap[confKey] = jsonValue
		} else {
			confMap[confKey] = confValue
		}
	}
	var app = new(resp.AppConfig)
	if fileUrl, ex := confMap[constant.FILE_URL]; ex {
		app.FileUrl = fileUrl.(string)
	}
	if resourceAuthority, ex := confMap[constant.RESOURCE_AUTHORITY]; ex {
		app.ResourceAuthority = resourceAuthority.(bool)
	}
	if shareUrl, ex := confMap[constant.SHARE_URL]; ex {
		app.ShareUrl = shareUrl.(string)
	}
	if downloadApp, ex := confMap[constant.DOWNLOAD_APP]; ex {
		app.DownloadApp = downloadApp.(string)
	}
	//格式化
	appJson, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	return appJson, nil
}
