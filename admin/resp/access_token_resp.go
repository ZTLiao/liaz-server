package resp

type AccessTokenResp struct {
	AccessToken string `json:"accessToken"` // 令牌
	ExpireAt    int64  `json:"expireAt"`    // 过期时间
	AdminId     int64  `json:"adminId"`     // 当前用户ID
}
