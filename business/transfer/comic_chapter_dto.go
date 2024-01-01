package transfer

type ComicChapterDto struct {
	ComicId        int64  `xorm:"comic_id"`
	Title          string `xorm:"title"`
	Cover          string `xorm:"cover"`
	ComicChapterId int64  `xorm:"comic_chapter_id"`
	ChapterName    string `xorm:"chapter_name"`
}
