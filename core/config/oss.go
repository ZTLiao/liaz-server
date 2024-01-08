package config

import (
	"core/system"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Oss struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

func (e *Oss) Init() {
	if e == nil {
		return
	}
	client, err := oss.New(e.Endpoint, e.AccessKeyId, e.AccessKeySecret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	system.SetOssClient(client)
}
