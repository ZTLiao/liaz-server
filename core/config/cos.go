package config

import (
	"core/system"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Cos struct {
	SecretId   string `yaml:"secretId"`
	SecretKey  string `yaml:"secretKey"`
	ServiceUrl string `yaml:"serviceUrl"`
	BucketUrl  string `yaml:"bucketUrl"`
}

func (e *Cos) Init() {
	if e == nil {
		return
	}
	bucketUrl, err := url.Parse(e.BucketUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	serviceUrl, err := url.Parse(e.ServiceUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	baseUrl := &cos.BaseURL{BucketURL: bucketUrl, ServiceURL: serviceUrl}
	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  e.SecretId,
			SecretKey: e.SecretKey,
		},
	})
	system.SetCosClient(client)
}
