package resp

import "core/types"

type NovelUpgradeResp struct {
	NovelChapterId int64      `json:"novelChapterId"`
	NovelId        int64      `json:"novelId"`
	Title          string     `json:"title"`
	Cover          string     `json:"cover"`
	Categories     []string   `json:"categories"`
	Authors        []string   `json:"authors"`
	UpgradeChapter string     `json:"upgradeChapter"`
	Updated        types.Time `json:"updated"`
}
