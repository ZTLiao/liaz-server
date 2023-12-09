package resp

import (
	"core/types"
)

type CustomOAuth2Token struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"token_type,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	Expiry       types.Time `json:"expiry,omitempty"`
	UserId       int64      `json:"user_id"`
}
