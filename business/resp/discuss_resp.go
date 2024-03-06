package resp

import "core/types"

type DiscussResp struct {
	DiscussId int64        `json:"discussId"`
	UserId    int64        `json:"userId"`
	Nickname  string       `json:"nickname"`
	Avatar    string       `json:"avatar"`
	Content   string       `json:"content"`
	CreatedAt types.Time   `json:"createdAt"`
	Paths     []string     `json:"paths"`
	Parent    *DiscussResp `json:"parent"`
}
