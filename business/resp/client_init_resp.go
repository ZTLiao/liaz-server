package resp

type ClientInitResp struct {
	Key *KeyConfig `json:"key"`
	App string     `json:"app"`
}

type KeyConfig struct {
	K1 string `json:"k1"`
	K2 string `json:"k2"`
}

type AppConfig struct {
	FileUrl           string                 `json:"fileUrl"`
	ResourceAuthority bool                   `json:"resourceAuthority"`
	ShareUrl          string                 `json:"shareUrl"`
	DownloadApp       string                 `json:"downloadApp"`
	Splash            map[string]interface{} `json:"splash"`
	EmptyPage         string                 `json:"emptyPage"`
	AdsFlag           int64                  `json:"adsFlag"`
}
