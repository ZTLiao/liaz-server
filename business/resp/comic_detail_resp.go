package resp

import (
	"core/types"
)

type ComicDetailResp struct {
	ComicId         int64             `json:"comicId"`
	Title           string            `json:"title"`
	Cover           string            `json:"cover"`
	AuthorIds       []int64           `json:"authorIds"`
	Authors         []string          `json:"authors"`
	CategoryIds     []int64           `json:"categoryIds"`
	Categories      []string          `json:"categories"`
	SubscribeNum    int32             `json:"subscribeNum"`
	HitNum          int32             `json:"hitNum"`
	Updated         types.Time        `json:"updated"`
	Description     string            `json:"description"`
	Flag            int8              `json:"flag"`
	Direction       int8              `json:"direction"`
	Volumes         []ComicVolumeResp `json:"volumes"`
	IsSubscribe     bool              `json:"isSubscribe"`
	BrowseChapterId int64             `json:"browseChapterId"`
	CurrentIndex    int               `json:"currentIndex"`
}
