package resp

import "core/types"

type ComicChapterResp struct {
	ComicChapterId int64      `json:"comicChapterId"`
	ComicId        int64      `json:"comicId"`
	Flag           int8       `json:"flag"`
	ChapterName    string     `json:"chapterName"`
	ChapterType    int8       `json:"chapterType"`
	PageNum        int        `json:"pageNum"`
	SeqNo          int        `json:"seqNo"`
	Direction      int8       `json:"direction"`
	UpdatedAt      types.Time `json:"updatedAt"`
	Paths          []string   `json:"paths"`
	CurrentIndex   int        `json:"currentIndex"`
}
