package transfer

import "core/types"

type ComicChapterDto struct {
	ComicId     int64      `json:"comicId" xorm:"comic_id"`
	Title       string     `json:"title" xorm:"title"`
	Cover       string     `json:"cover" xorm:"cover"`
	EndTime     types.Time `json:"endTime" xorm:"end_time"`
	ChapterId   int64      `json:"chapterId" xorm:"chapter_id"`
	ChapterName string     `json:"chapterName" xorm:"chapter_name"`
	IsUpgrade   int8       `json:"isUpgrade" xorm:"is_upgrade"`
}
