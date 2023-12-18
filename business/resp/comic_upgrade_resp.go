package resp

import "core/types"

type ComicUpgradeResp struct {
	ComicChapterId int64      `json:"comicChapterId"`
	ComicId        int64      `json:"comicId"`
	Title          string     `json:"title"`
	Cover          string     `json:"cover"`
	Categories     []string   `json:"categories"`
	Authors        []string   `json:"authors"`
	UpgradeChapter string     `json:"upgradeChapter"`
	Updated        types.Time `json:"updated"`
}
