package transfer

import "core/types"

type ComicChapterDto struct {
	ComicId     int64      `xorm:"comic_id"`
	Title       string     `xorm:"title"`
	Cover       string     `xorm:"cover"`
	EndTime     types.Time `xorm:"end_time"`
	ChapterId   int64      `xorm:"chapter_id"`
	ChapterName string     `xorm:"chapter_name"`
}
