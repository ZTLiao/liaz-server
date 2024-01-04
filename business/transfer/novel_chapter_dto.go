package transfer

import "core/types"

type NovelChapterDto struct {
	NovelId     int64      `xorm:"novel_id"`
	Title       string     `xorm:"title"`
	Cover       string     `xorm:"cover"`
	EndTime     types.Time `xorm:"end_time"`
	ChapterId   int64      `xorm:"chapter_id"`
	ChapterName string     `xorm:"chapter_name"`
	IsUpgrade   int8       `xorm:"is_upgrade"`
}
