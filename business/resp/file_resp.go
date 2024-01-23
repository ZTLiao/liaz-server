package resp

type FileResp struct {
	Path       string `json:"path"`
	ExpireTime int64  `json:"expireTime"`
	RequestUri string `json:"requestUri"`
}
