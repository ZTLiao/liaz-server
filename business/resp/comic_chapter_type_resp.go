package resp

type ComicChapterTypeResp struct {
	ChapterType int8               `json:"chapterType"`
	Flag        int8               `json:"flag"`
	Chapters    []ComicChapterResp `json:"chapters"`
}
