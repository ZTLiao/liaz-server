package resp

type ComicChapterResp struct {
	ComicChapterId int64    `json:"comicChapterId"`
	ComicId        int64    `json:"comicId"`
	Flag           int8     `json:"flag"`
	ChapterName    string   `json:"chapterName"`
	SeqNo          int      `json:"seqNo"`
	Paths          []string `json:"paths"`
	Direction      int8     `json:"direction"`
}
