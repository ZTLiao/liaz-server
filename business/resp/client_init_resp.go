package resp

type ClientInitResp struct {
	Key *KeyConfig `json:"key"`
	App string     `json:"config"`
}

type KeyConfig struct {
	K1 string `json:"k1"`
}

type AppConfig struct {
	FileUrl string `json:"fileUrl"`
}
