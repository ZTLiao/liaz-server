package resp

import "core/types"

type NovelChapterResp struct {
	NovelChapterId int64      `json:"novelChapterId"`
	NovelId        int64      `json:"novelId"`
	Flag           int8       `json:"flag"`
	ChapterName    string     `json:"chapterName"`
	ChapterType    int8       `json:"chapterType"`
	PageNum        int        `json:"pageNum"`
	SeqNo          int        `json:"seqNo"`
	Direction      int8       `json:"direction"`
	UpdatedAt      types.Time `json:"updatedAt"`
	Paths          []string   `json:"paths"`
	Types          []string   `json:"types"`
}
