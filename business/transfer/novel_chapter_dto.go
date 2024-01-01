package transfer

type NovelChapterDto struct {
	NovelId        int64  `xorm:"novel_id"`
	Title          string `xorm:"title"`
	Cover          string `xorm:"cover"`
	NovelChapterId int64  `xorm:"novel_chapter_id"`
	ChapterName    string `xorm:"chapter_name"`
}
