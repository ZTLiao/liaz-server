package config

import (
	"core/system"
	"core/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
)

const (
	DEV  string = "dev"
	TEST string = "test"
	PROD string = "prod"
)

var (
	iConfigClient config_client.IConfigClient
	profiles      map[string]Nacos = map[string]Nacos{
		DEV: {
			Username:           "nacos",
			Password:           "nacos",
			ServerAddr:         "127.0.0.1:8848",
			Namespace:          "5f49cd28-4eb2-4b0b-a467-bab25f4c9535",
			SharedDataIds:      "application.yaml,database.yaml",
			RefreshableDataIds: "application.yaml",
		},
		TEST: {
			Username:           "nacos",
			Password:           "nacos",
			ServerAddr:         "127.0.0.1:8848",
			Namespace:          "e85ee8e9-472c-4803-bcf0-0b0f7d0ea0b7",
			SharedDataIds:      "application.yaml,database.yaml",
			RefreshableDataIds: "application.yaml",
		},
		PROD: {
			Username:           "nacos",
			Password:           "nacos",
			ServerAddr:         "172.17.0.1:8848",
			Namespace:          "b8a5a983-8632-40f4-9e83-6783a4b1680c",
			SharedDataIds:      "application.yaml,database.yaml",
			RefreshableDataIds: "application.yaml",
		},
	}
)

type Nacos struct {
	Username           string
	Password           string
	ServerAddr         string
	Namespace          string
	SharedDataIds      string
	RefreshableDataIds string
}

func (e *Nacos) Init() {
	//nacos地址
	serverAddrArray := strings.Split(e.ServerAddr, utils.COLON)
	ipAddr := serverAddrArray[0]
	port, err := strconv.Atoi(serverAddrArray[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ipAddr, uint64(port), constant.WithContextPath("/nacos")),
	}
	cc := *constant.NewClientConfig(
		constant.WithUsername(e.Username),
		constant.WithPassword(e.Password),
		constant.WithNamespaceId(e.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithLogDir("logs"),
		constant.WithCacheDir("cache"),
		constant.WithLogLevel("debug"),
	)
	iClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	iConfigClient = iClient
	system.SetIConfigClient(iClient)
	//获取配置
	e.config()
	//设置监听事件
	e.listen()
}

// 获取配置
func (e *Nacos) config() {
	var configYml string
	sharedDataIdArray := strings.Split(e.SharedDataIds, utils.COMMA)
	for i := 0; i < len(sharedDataIdArray); i++ {
		dataId := sharedDataIdArray[i]
		fmt.Printf("Loading dataId : [%s]\n", dataId)
		content, err := iConfigClient.GetConfig(
			vo.ConfigParam{
				DataId: dataId,
			},
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		configYml += fmt.Sprintf("%s\n", content)
	}
	//当前应用的yaml优先级最高
	content, err := iConfigClient.GetConfig(
		vo.ConfigParam{
			DataId: system.GetName() + ".yaml",
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	configYml += fmt.Sprintf("%s\n", content)
	//刷新全局配置
	e.refresh(configYml)
}

// 设置监听事件
func (e *Nacos) listen() {
	refreshableDataIdArray := strings.Split(e.RefreshableDataIds, utils.COMMA)
	for i := 0; i < len(refreshableDataIdArray); i++ {
		dataId := refreshableDataIdArray[i]
		err := iConfigClient.ListenConfig(
			vo.ConfigParam{
				DataId: dataId,
				Group:  "DEFAULT_GROUP",
				OnChange: func(namespace, group, dataId, data string) {
					fmt.Printf("Listening config group : %s, dataId : %s, data : %s\n", group, dataId, data)
					e.refresh(data)
				},
			},
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

// 刷新全局配置
func (e *Nacos) refresh(configYml string) {
	if err := yaml.Unmarshal([]byte(configYml), &SystemConfig); err != nil {
		fmt.Println(err.Error())
	}
}
