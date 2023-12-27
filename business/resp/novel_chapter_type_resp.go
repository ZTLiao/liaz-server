package resp

type NovelChapterTypeResp struct {
	ChapterType int8               `json:"chapterType"`
	Flag        int8               `json:"flag"`
	Chapters    []NovelChapterResp `json:"chapters"`
}
