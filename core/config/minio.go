package config

import (
	"core/system"
	"core/utils"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyId"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Secure          bool   `yaml:"secure"`
}

func (e *Minio) Init() {
	if e == nil {
		return
	}
	client, err := minio.New(e.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(e.AccessKeyID, e.SecretAccessKey, utils.EMPTY),
		Secure: e.Secure,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	system.SetMinioClient(client)
}
