package config

import (
	"core/constant"
	"core/system"

	"golang.org/x/oauth2"
)

type Oauth2 struct {
	ClientId      string `yaml:"clientId"`
	ClientSecret  string `yaml:"clientSecret"`
	AuthServerUrl string `yaml:"authServerUrl"`
}

func (e *Oauth2) Init() {
	if e == nil {
		return
	}
	system.SetOauth2Config(&oauth2.Config{
		ClientID:     e.ClientId,
		ClientSecret: e.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: e.AuthServerUrl + constant.OAUTH_TOKEN,
		},
	})
}
