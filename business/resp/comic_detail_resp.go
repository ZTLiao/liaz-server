package resp

import (
	"core/types"
)

type ComicDetailResp struct {
	ComicId      int64              `json:"comicId"`
	Title        string             `json:"title"`
	Cover        string             `json:"cover"`
	AuthorIds    []int64            `json:"authorIds"`
	Authors      []string           `json:"authors"`
	CategoryIds  []int64            `json:"categoryIds"`
	Categories   []string           `json:"categories"`
	SubscribeNum int32              `json:"subscribeNum"`
	HitNum       int32              `json:"hitNum"`
	Updated      types.Time         `json:"updated"`
	Flag         int8               `json:"flag"`
	Direction    int8               `json:"direction"`
	Chapters     []ComicChapterResp `json:"chapters"`
}
